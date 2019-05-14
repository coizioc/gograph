package main

// DistQueue is a priority queue sorted by distance.
type DistQueue struct {
	dlist []*vertexDist
}

type vertexDist struct {
	vertex int
	dist   int
}

// NewDistQueue initializes a new DistQueue.
func NewDistQueue() *DistQueue {
	return &DistQueue{[]*vertexDist{}}
}

// Enqueue enqueues an element from DistQueue and sorts it.
func (d *DistQueue) Enqueue(v, dist int) {
	d.dlist = append(d.dlist, &vertexDist{v, dist})
	d.sort()
}

// Dequeue dequeues an element from DistQueue.
func (d *DistQueue) Dequeue() (int, int) {
	ret := d.dlist[len(d.dlist)-1]
	d.dlist = d.dlist[:len(d.dlist)-1]
	return ret.vertex, ret.dist
}

// Size returns the size of the DistQueue.
func (d *DistQueue) Size() int {
	return len(d.dlist)
}

// Update updates the distance for a given vertex in DistQueue.
func (d *DistQueue) Update(v, newDist int) {
	for i := range d.dlist {
		if d.dlist[i].vertex == v {
			d.dlist[i].dist = newDist
			break
		}
	}
	d.sort()
}

func (d *DistQueue) sort() {
	for i := len(d.dlist) - 1; i > 0; i-- {
		if d.dlist[i].dist > d.dlist[i-1].dist {
			d.dlist[i], d.dlist[i-1] = d.dlist[i-1], d.dlist[i]
		}
	}
}
