module Gelm where


import Json.Decode as JD exposing ((:=))
import Json.Encode as JE


type Enum
  = EnumValueDefault -- 0
  | EnumValue1 -- 1
  | EnumValue2 -- 2


enumDecoder : JD.Decoder Enum
enumDecoder =
  let
    lookup s = case s of
      "ENUM_VALUE_DEFAULT" -> EnumValueDefault
      "ENUM_VALUE_1" -> EnumValue1
      "ENUM_VALUE_2" -> EnumValue2
      _ -> EnumValueDefault
  in
    JD.map lookup JD.string


enumEncoder : Enum -> JE.Value
enumEncoder v =
  let
    lookup s = case s of
      EnumValueDefault -> "ENUM_VALUE_DEFAULT"
      EnumValue1 -> "ENUM_VALUE_1"
      EnumValue2 -> "ENUM_VALUE_2"
  in
    JE.string <| lookup v


type alias SubMessage =
  { id : Int
  }


subMessageDecoder : JD.Decoder SubMessage
subMessageDecoder =
  JD.object1 SubMessage
    ("id" := JD.int)


subMessageEncoder : SubMessage -> JE.Value
subMessageEncoder v =
  JE.object
    [ ("id", JE.int v.id)
    ]


type alias Message =
  { id : Int
  , fieldWithLongName : String
  , enum : Enum
  , subMessage : SubMessage
  }


messageDecoder : JD.Decoder Message
messageDecoder =
  JD.object4 Message
    ("id" := JD.int)
    ("fieldWithLongName" := JD.string)
    ("enum" := enumDecoder)
    ("subMessage" := subMessageDecoder)


messageEncoder : Message -> JE.Value
messageEncoder v =
  JE.object
    [ ("id", JE.int v.id)
    , ("fieldWithLongName", JE.string v.fieldWithLongName)
    , ("enum", enumEncoder v.enum)
    , ("subMessage", subMessageEncoder v.subMessage)
    ]


