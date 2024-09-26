package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "net/http"
)

func (g *GoogleDevice) Sync(checkin *GoogleCheckin) error {
   field_7, err := checkin.field_7()
   if err != nil {
      return err
   }
   message := protobuf.Message{}
   message.AddFunc(1, func(m protobuf.Message) {
      m.AddFunc(10, func(m protobuf.Message) {
         for _, value := range g.Feature {
            m.AddFunc(1, func(m protobuf.Message) {
               m.AddBytes(1, []byte(value))
            })
         }
         for _, value := range g.Library {
            m.AddBytes(2, []byte(value))
         }
         for _, value := range g.Texture {
            m.AddBytes(4, []byte(value))
         }
      })
   })
   message.AddFunc(1, func(m protobuf.Message) {
      m.AddFunc(15, func(m protobuf.Message) {
         m.AddBytes(4, []byte(g.Abi))
      })
   })
   message.AddFunc(1, func(m protobuf.Message) {
      m.AddFunc(18, func(m protobuf.Message) {
         m.AddBytes(1, []byte("am-unknown")) // X-DFE-Client-Id
      })
   })
   message.AddFunc(1, func(m protobuf.Message) {
      m.AddFunc(19, func(m protobuf.Message) {
         m.AddVarint(2, google_play_store)
      })
   })
   message.AddFunc(1, func(m protobuf.Message) {
      m.AddFunc(21, func(m protobuf.Message) {
         m.AddVarint(6, gl_es_version)
      })
   })
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/sync",
      bytes.NewReader(message.Marshal()),
   )
   if err != nil {
      return err
   }
   x_dfe_device_id(req, field_7)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   return nil
}
