package domain

type ServerList struct {
    ServerItems []ServerItem
}

func NewServerList() *ServerList {
    return &ServerList {
        ServerItems: make([]ServerItem, 0),
    }
}

func (sl *ServerList) AddServer(name string, addr string, port int) {
    var newServer ServerItem = ServerItem {
        Name: name,
        Addr: addr,
        Port: port,
    }
    sl.ServerItems = append(sl.ServerItems, newServer)
}

