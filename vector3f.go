package vector

import "math32"
import "math"

type Vector3f struct {
	X, Y, Z float32
}

func SliceToVector3f(slice []float32) Vector3f {
	return Vector3f{X:slice[0],Y:slice[1],Z:slice[2]}
}

func FloatToVector3f(n float32) Vector3f {
	return Vector3f{n,n,n}
}

func (self *Vector3f) ToSlice() []float32 {
	return []float32{self.X, self.Y, self.Z}
}

func (self *Vector3f) Iter() <-chan float32 {
	ch := make(chan float32)
	go func() {
		ch <- self.X
		ch <- self.Y
		ch <- self.Z
	}()
	return ch
}

func (self *Vector3f) XYZ(i int) float32 {
	var retval float32
	switch i {
	case 0:
		retval = self.X
	case 1:
		retval = self.Y
	case 2:
		retval = self.Z
	default:
		panic("illegal index number")
	}
	return retval
}

func (self *Vector3f) Copy() Vector3f {
	return Vector3f{self.X, self.Y, self.Z}
}

func (self *Vector3f) Add(other Vector3f) Vector3f {
	return Vector3f{self.X + other.X,
		self.Y + other.Y,
		self.Z + other.Z}
}

func (self *Vector3f) Sub(other Vector3f) Vector3f {
	return Vector3f{self.X - other.X,
		self.Y - other.Y,
		self.Z - other.Z}
}

func (self *Vector3f) Mulf(v float32) Vector3f {
	return Vector3f{self.X * v, self.Y * v, self.Z * v}
}

func (self Vector3f) Mulv(other Vector3f) Vector3f {
	return Vector3f{self.X * other.X, self.Y * other.Y, self.Z * other.Z}
}

func (self *Vector3f) Neg() {
	self.X *= -1
	self.Y *= -1
	self.Z *= -1
}

func (self *Vector3f) Dot(other Vector3f) float32 {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func (self *Vector3f) Cross(other Vector3f) Vector3f {
	return Vector3f{(self.Y * other.Z) - (self.Z * other.Y),
		(self.Z * other.X) - (self.X * other.Z),
		(self.X * other.Y) - (self.Y * other.X)}
}

func (self *Vector3f) Is_Zero() bool {
	if self.X == 0.0 && self.Y == 0 && self.Z == 0.0 {
		return true
	}
	return false
}

func (self *Vector3f) Mag() float32 {
	return math32.Sqrt(self.X*self.X + self.Y*self.Y + self.Z*self.Z)
}

func (self *Vector3f) Unitize() {
	/* Zero vectors, and vectors of near zero magnitude, produce zero length,
      and (since 1 / 0 is conditioned to 0) ultimately a zero vector result.
      Vectors of extremely large magnitude produce +infinity length, and (since
      1 / inf is 0) ultimately a zero vector result.
      (Perhaps zero vectors should produce infinite results, but pragmatically,
      zeros are probably easier to handle than infinities.) */
	var length_inv float32
	length := self.Mag()
	if length != 0 {
		length_inv = 1.0 / length
	} else {
		length_inv = 0.0
	}
	self.X *= length_inv
	self.Y *= length_inv
	self.Z *= length_inv
}

func (self Vector3f) UnitizeCopy() Vector3f {
	v := self.Copy()
	v.Unitize()
	return v
}

func (self *Vector3f) Clamped(lo, hi Vector3f) Vector3f {
	//clamp value to be within vectors lo and hi
	return Vector3f{math32.Min(math32.Max(self.X, lo.X), hi.X),
		math32.Min(math32.Max(self.Y, lo.Y), hi.Y),
		math32.Min(math32.Max(self.Z, lo.Z), hi.Z)}
}

var VECTOR_ZERO = Vector3f{0.0, 0.0, 0.0}
var VECTOR_ONE = Vector3f{1.0, 1.0, 1.0}
var VECTOR_MAX = Vector3f{math.MaxFloat32,
	math.MaxFloat32,
	math.MaxFloat32}
