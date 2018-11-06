package main
import (
    "syscall"
    "fmt"
    "os"
    "unsafe"
)

func decode(buf []byte) int {
    e := *(*syscall.InotifyEvent)(unsafe.Pointer(&buf[0]))
    if (e.Mask & syscall.IN_ACCESS != 0) {
        fmt.Printf("IN_ACCESS ")
    }
    if (e.Mask & syscall.IN_ATTRIB != 0) {
        fmt.Printf("IN_ATTRIB ")
    }
    if (e.Mask & syscall.IN_CLOSE_NOWRITE != 0) {
        fmt.Printf("IN_CLOSE_NOWRITE ")
    }
    if (e.Mask & syscall.IN_CLOSE_WRITE != 0) {
        fmt.Printf("IN_CLOSE_WRITE ")
    }
    if (e.Mask & syscall.IN_CREATE != 0) {
        fmt.Printf("IN_CREATE ")
    }
    if (e.Mask & syscall.IN_DELETE != 0) {
        fmt.Printf("IN_DELETE ")
    }
    if (e.Mask & syscall.IN_DELETE_SELF != 0) {
        fmt.Printf("IN_DELETE_SELF ")
    }
    if (e.Mask & syscall.IN_IGNORED != 0) {
        fmt.Printf("IN_IGNORED ")
    }
    if (e.Mask & syscall.IN_ISDIR != 0) {
        fmt.Printf("IN_ISDIR ")
    }
    if (e.Mask & syscall.IN_MODIFY != 0) {
        fmt.Printf("IN_MODIFY ")
    }
    if (e.Mask & syscall.IN_MOVE_SELF != 0) {
        fmt.Printf("IN_MOVE_SELF ")
    }
    if (e.Mask & syscall.IN_MOVED_FROM != 0) {
        fmt.Printf("IN_MOVED_FROM ")
    }
    if (e.Mask & syscall.IN_MOVED_TO != 0) {
        fmt.Printf("IN_MOVED_TO ")
    }
    if (e.Mask & syscall.IN_OPEN != 0) {
        fmt.Printf("IN_OPEN ")
    }
    if (e.Mask & syscall.IN_Q_OVERFLOW != 0) {
        fmt.Printf("IN_Q_OVERFLOW ")
    }
    if (e.Mask & syscall.IN_UNMOUNT != 0) {
        fmt.Printf("IN_UNMOUNT ")
    }
    name := string(buf[syscall.SizeofInotifyEvent:syscall.SizeofInotifyEvent+e.Len])
    fmt.Println(name)

    return syscall.SizeofInotifyEvent+int(e.Len)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("usage: ./inotify file-to-be-monitored")
        return
    }

    file := os.Args[1]

    const Size = 10*(syscall.SizeofInotifyEvent+syscall.NAME_MAX+1)
    buf := make([]byte, Size)

    fd, err := syscall.InotifyInit()

    if err != nil {
        fmt.Println(err)
        return
    }

    _, err = syscall.InotifyAddWatch(fd, file, syscall.IN_ALL_EVENTS)

    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        n, err := syscall.Read(fd, buf)
        if err != nil {
            fmt.Println(err)
            return
        }

        for offset := 0; offset < n; {
            eventLen := decode(buf[offset:])
            offset += eventLen
        }

    }
}
