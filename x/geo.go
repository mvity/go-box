/*
 * Copyright © 2021 - 2023 vity <vityme@icloud.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package x

import (
	"bytes"
	"math"
	"strings"
)

type GeoPoint struct {
	Lng float64
	Lat float64
}

// NewGeoPoint 生成GeoPoint对象
func NewGeoPoint(latLng string) GeoPoint {
	strs := strings.Split(latLng, ",")
	return GeoPoint{
		Lng: ToFloat64(strs[0]),
		Lat: ToFloat64(strs[1]),
	}
}

// NewGeoPoints 生成GeoPoint对象数组
func NewGeoPoints(latLngs string) []GeoPoint {
	points := make([]GeoPoint, 0)
	for _, s := range strings.Split(latLngs, ":") {
		points = append(points, NewGeoPoint(s))
	}
	return points
}

// GeoCheckInAreas 检查指定点是否在区域内
func GeoCheckInAreas(point GeoPoint, areas []GeoPoint) bool {
	if len(areas) < 3 {
		return false
	}
	sPoint := areas[0]
	ePoint := areas[len(areas)-1]
	if sPoint.Lng != ePoint.Lng && sPoint.Lat != ePoint.Lat {
		areas = append(areas, sPoint)
	}

	pointNum := len(areas) //点个数
	intersectCount := 0    //cross points count of x
	precision := 2e-10     //浮点类型计算时候与0比较时候的容差
	p1 := GeoPoint{}       //neighbour bound vertices
	p2 := GeoPoint{}
	p := point //测试点

	p1 = areas[0] //left vertex
	for i := 0; i < pointNum; i++ {
		if p.Lng == p1.Lng && p.Lat == p1.Lat {
			return true
		}
		p2 = areas[i%pointNum]
		if p.Lat < math.Min(p1.Lat, p2.Lat) || p.Lat > math.Max(p1.Lat, p2.Lat) {
			p1 = p2
			continue //next ray left point
		}

		if p.Lat > math.Min(p1.Lat, p2.Lat) && p.Lat < math.Max(p1.Lat, p2.Lat) {
			if p.Lng <= math.Max(p1.Lng, p2.Lng) { //x is before of ray
				if p1.Lat == p2.Lat && p.Lng >= math.Min(p1.Lng, p2.Lng) {
					return true
				}

				if p1.Lng == p2.Lng { //ray is vertical
					if p1.Lng == p.Lng { //overlies on a vertical ray
						return true
					} else { //before ray
						intersectCount++
					}
				} else { //cross point on the left side
					xinters := (p.Lat-p1.Lat)*(p2.Lng-p1.Lng)/(p2.Lat-p1.Lat) + p1.Lng
					if math.Abs(p.Lng-xinters) < precision {
						return true
					}

					if p.Lng < xinters { //before ray
						intersectCount++
					}
				}
			}
		} else { //special case when ray is crossing through the vertex
			if p.Lat == p2.Lat && p.Lng <= p2.Lng { //p crossing over p2
				p3 := areas[(i+1)%pointNum]
				if p.Lat >= math.Min(p1.Lat, p3.Lat) && p.Lat <= math.Max(p1.Lat, p3.Lat) {
					intersectCount++
				} else {
					intersectCount += 2
				}
			}
		}
		p1 = p2 //next ray left point
	}
	if intersectCount%2 == 0 { //偶数在多边形外
		return false
	} else { //奇数在多边形内
		return true
	}
}

const (
	geoHashbase32               = "0123456789bcdefghjkmnpqrstuvwxyz"
	geoHashmaxLatitude  float64 = 90
	geoHashminLatitude  float64 = -90
	geoHashmaxLongitude float64 = 180
	geoHashminLongitude float64 = -180
)

var (
	geoHashbits        = []int{16, 8, 4, 2, 1}
	geoHashbase32Bytes = []byte(geoHashbase32)
)

// geoHashBox geohash的精度与其长度成正比
// 每个点的geohash值实际上代表了一个区域，这个区域的大小与geohash的精度成反比
// 坐标点的格式为（纬度，经度）
// 将这个区域用一个矩形表示
type geoHashBox struct {
	MinLat, MaxLat float64 // 纬度
	MinLng, MaxLng float64 // 经度
}

func (b *geoHashBox) Width() float64 {
	return b.MaxLng - b.MinLng
}

func (b *geoHashBox) Height() float64 {
	return b.MaxLat - b.MinLat
}

// geohash精度的设定参考 http://en.wikipedia.org/wiki/Geohash
// geohash length	lat geoHashbits	lng geoHashbits	lat error	lng error	km error
// 1				2			3			±23			±23			±2500
// 2				5			5			± 2.8		± 5.6		±630
// 3				7			8			± 0.70		± 0.7		±78
// 4				10			10			± 0.087		± 0.18		±20
// 5				12			13			± 0.022		± 0.022		±2.4
// 6				15			15			± 0.0027	± 0.0055	±0.61
// 7				17			18			±0.00068	±0.00068	±0.076
// 8				20			20			±0.000085	±0.00017	±0.019

// geoHashEncode 输入值：纬度，经度，精度(geohash的长度),返回geohash, 以及该点所在的区域
func geoHashEncode(latitude float64, longitude float64, precision int) (string, *geoHashBox) {
	var geohash bytes.Buffer
	var minLat, maxLat = geoHashminLatitude, geoHashmaxLatitude
	var minLng, maxLng = geoHashminLongitude, geoHashmaxLongitude
	var mid float64 = 0

	bit, ch, length, isEven := 0, 0, 0, true
	for length < precision {
		if isEven {
			if mid = (minLng + maxLng) / 2; mid < longitude {
				ch |= geoHashbits[bit]
				minLng = mid
			} else {
				maxLng = mid
			}
		} else {
			if mid = (minLat + maxLat) / 2; mid < latitude {
				ch |= geoHashbits[bit]
				minLat = mid
			} else {
				maxLat = mid
			}
		}

		isEven = !isEven
		if bit < 4 {
			bit++
		} else {
			geohash.WriteByte(geoHashbase32Bytes[ch])
			length, bit, ch = length+1, 0, 0
		}
	}

	b := &geoHashBox{
		MinLat: minLat,
		MaxLat: maxLat,
		MinLng: minLng,
		MaxLng: maxLng,
	}

	return geohash.String(), b
}

// GeoHashEncode 输入值：纬度，经度，精度(geohash的长度)
func GeoHashEncode(longitude float64, latitude float64, precision int) string {
	val, _ := geoHashEncode(latitude, longitude, precision)
	return val
}

// GeoHashGetNeighbors 计算该点（latitude, longitude）在精度precision下的邻居 -- 周围8个区域+本身所在区域 返回这些区域的geohash值，总共9个
func GeoHashGetNeighbors(latitude, longitude float64, precision int) []string {
	geohashs := make([]string, 9)

	// 本身
	geohash, b := geoHashEncode(latitude, longitude, precision)
	geohashs[0] = geohash

	// 上下左右
	geohashUp, _ := geoHashEncode((b.MinLat+b.MaxLat)/2+b.Height(), (b.MinLng+b.MaxLng)/2, precision)
	geohashDown, _ := geoHashEncode((b.MinLat+b.MaxLat)/2-b.Height(), (b.MinLng+b.MaxLng)/2, precision)
	geohashLeft, _ := geoHashEncode((b.MinLat+b.MaxLat)/2, (b.MinLng+b.MaxLng)/2-b.Width(), precision)
	geohashRight, _ := geoHashEncode((b.MinLat+b.MaxLat)/2, (b.MinLng+b.MaxLng)/2+b.Width(), precision)

	// 四个角
	geohashLeftUp, _ := geoHashEncode((b.MinLat+b.MaxLat)/2+b.Height(), (b.MinLng+b.MaxLng)/2-b.Width(), precision)
	geohashLeftDown, _ := geoHashEncode((b.MinLat+b.MaxLat)/2-b.Height(), (b.MinLng+b.MaxLng)/2-b.Width(), precision)
	geohashRightUp, _ := geoHashEncode((b.MinLat+b.MaxLat)/2+b.Height(), (b.MinLng+b.MaxLng)/2+b.Width(), precision)
	geohashRightDown, _ := geoHashEncode((b.MinLat+b.MaxLat)/2-b.Height(), (b.MinLng+b.MaxLng)/2+b.Width(), precision)

	geohashs[1], geohashs[2], geohashs[3], geohashs[4] = geohashUp, geohashDown, geohashLeft, geohashRight
	geohashs[5], geohashs[6], geohashs[7], geohashs[8] = geohashLeftUp, geohashLeftDown, geohashRightUp, geohashRightDown

	return geohashs
}
