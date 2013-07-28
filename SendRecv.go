package main

// Flags for cgo obtained from OpenMPI using:
//      $ mpicc --showme:compile
//      $ mpicc --showme:link

// #cgo CFLAGS: -I/usr/local/Cellar/open-mpi/1.4.5/include
// #cgo LDFLAGS: -L/usr/local/Cellar/open-mpi/1.4.5/lib -lmpi -lopen-rte -lopen-pal -lutil
// #include <mpi.h>
// 
// MPI_Comm get_world() {
//      return (MPI_Comm)(MPI_COMM_WORLD);
// }
// MPI_Datatype get_mpi_int() {
//      return (MPI_Datatype)(MPI_INT);
// }
import "C"

import (
    "fmt"
    "log"
    "unsafe"
)

func main() {
    err := C.MPI_Init(nil, nil)
    if err != 0 {
        log.Fatal(err)
    }

    comm := C.get_world()
    rank := C.int(-1)
    C.MPI_Comm_rank(comm, &rank)

    buf := C.int(-1)

    if int(rank) == 0 {
        buf = C.int(42)
        C.MPI_Send(unsafe.Pointer(&buf), 1, C.get_mpi_int(), 1, 1, comm)
    } else if int(rank) == 1 {
        C.MPI_Recv(unsafe.Pointer(&buf), 1, C.get_mpi_int(), 0, 1, comm,nil)
        fmt.Println("Rank = ", rank, " received value ", int(buf))
    }

    err = C.MPI_Finalize()
    if err != 0 {
        log.Fatal(err)
    }

}
