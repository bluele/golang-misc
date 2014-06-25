package ws

type Connection interface {
  ReadPump()
  Write(mt int, payload []byte) error 
  WritePump()
}