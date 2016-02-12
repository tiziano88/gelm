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


type alias Message =
  { id : Int
  , fieldWithLongName : String
  , enum : Enum
  }


messageDecoder : JD.Decoder Message
messageDecoder =
  JD.object3 Message
    ("id" := JD.int)
    ("fieldWithLongName" := JD.string)
    ("enum" := enumDecoder)


messageEncoder : Message -> JE.Value
messageEncoder v =
  JE.object
    [ ("id", JE.int v.id)
    , ("fieldWithLongName", JE.string v.fieldWithLongName)
    , ("enum", enumEncoder v.enum)
    ]


