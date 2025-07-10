package main

type Ray struct {
	Origin, Direction Vec3
}

// Methods

func (r *Ray) PointAt(t float64) Vec3 {
	return Add(r.Origin, Scale(r.Direction, t))
}
