package main

// Flags for cgo obtained from OpenMPI using:
//      $ mpicc --showme:compile
//      $ mpicc --showme:link

// #cgo CFLAGS: -I/usr/local/Cellar/open-mpi/1.4.5/include
// #cgo LDFLAGS: -L/usr/local/Cellar/open-mpi/1.4.5/lib -lmpi -lopen-rte -lopen-pal -lutil
// #include <mpi.h>
// 
// MPI_Comm get_MPI_COMM_WORLD() {
//      return (MPI_Comm)(MPI_COMM_WORLD);
// }
import "C"

import (
    "fmt"
    "log"
)

func main() {
    err := C.MPI_Init(nil, nil)
    if err != 0 {
        log.Fatal(err)
    }
    comm := C.get_MPI_COMM_WORLD()

    rank := C.int(-1)
    C.MPI_Comm_rank(comm,&rank)

    size := C.int(-1)
    C.MPI_Comm_size(comm,&size)

    fmt.Println("Hello from Go! Rank is ", rank, " of ", size)

    err = C.MPI_Finalize()
    if err != 0 {
        log.Fatal(err)
    }
}
