//share reasource within a pool
package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	ErrPoolClosed = errors.New("Current resources pool has been closed!")
)

type Pool struct {
	//define a resource type
	factory func()(io.Closer,error) //the resource can be use
	resources chan io.Closer //get the resource from channel
	mtx sync.Mutex //A lock when getting resource
	closed bool //is the pool close
}

//consstruct the pool with a factory and the size of the pool
func New(factory func() (io.Closer,error),size int) (*Pool,error){
	if size <= 0{
		return nil,errors.New("invalid size")
	}

	//return pool pointer and no errpr
	return &Pool{
		factory: factory, //each
		resources: make(chan io.Closer,size),
		closed: false,
	},nil
}

//required resource
func(p *Pool)AcquireResource() (io.Closer,error){
	//get the io.Closer resource and some error
	select {
	case resource,ok:= <- p.resources: //can get resource? or close
		if !ok{
			//there are no any resource
			return nil,ErrPoolClosed
		}
		fmt.Println("Acquired resource from the pool")
		return resource,nil
	default:
		fmt.Println("New resource from factory")
		return p.factory() //DBConnection is the resource later add to resource
	}
}

//release the resource
func(p *Pool)ReleaseResource(resource io.Closer){
	//check the resource pool
	p.mtx.Lock()
	defer p.mtx.Unlock()

	//check the pool is closed?
	if p.closed{
		resource.Close() //closed the resource
		return
	}

	select {
	case p.resources <- resource:
		fmt.Println("Release resource is putting back to the pool")
	default:
		fmt.Println("Pool is now full, resource is closing")
		resource.Close()
	}
}

func(p* Pool) Close(){
	//protect close variable,case race condition
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		fmt.Println("Pool is already Closed")
	}

	p.closed = true
	close(p.resources)
	for resource := range p.resources{
		//close all resource
		resource.Close()
	}
}
