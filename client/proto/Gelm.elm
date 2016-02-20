module Gelm where


import Json.Decode as JD exposing ((:=))
import Json.Encode as JE


optional : JD.Decoder a -> JD.Decoder (Maybe a)
optional decoder =
  JD.oneOf
    [ JD.map Just decoder
    , JD.succeed Nothing
    ]


withDefault : a -> JD.Decoder a -> JD.Decoder a
withDefault default decoder =
  JD.oneOf
    [ decoder
    , JD.succeed default
    ]


intField : String -> JD.Decoder Int
intField name =
  withDefault 0 (name := JD.int)


boolField : String -> JD.Decoder Bool
boolField name =
  withDefault False (name := JD.bool)


stringField : String -> JD.Decoder String
stringField name =
  withDefault "" (name := JD.string)


messageField : JD.Decoder a -> String -> JD.Decoder (Maybe a)
messageField decoder name =
  optional (name := decoder)


enumField : JD.Decoder a -> String -> JD.Decoder a
enumField decoder name =
  (name := decoder)


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
    (intField "id")


subMessageEncoder : SubMessage -> JE.Value
subMessageEncoder v =
  JE.object
    [ ("id", JE.int v.id)
    ]


type alias Message =
  { id : Int
  , fieldWithLongName : String
  , enum : Enum
  , subMessage : Maybe SubMessage
  , boolField : Bool
  , user : Maybe User
  }


messageDecoder : JD.Decoder Message
messageDecoder =
  JD.object6 Message
    (intField "id")
    (stringField "fieldWithLongName")
    (enumField enumDecoder "enum")
    (messageField subMessageDecoder "subMessage")
    (boolField "boolField")
    (messageField userDecoder "user")


messageEncoder : Message -> JE.Value
messageEncoder v =
  JE.object
    [ ("id", JE.int v.id)
    , ("fieldWithLongName", JE.string v.fieldWithLongName)
    , ("enum", enumEncoder v.enum)
    , ("boolField", JE.bool v.boolField)
    ]


type alias Address =
  { line1 : String
  , line2 : String
  , city : String
  , country : String
  }


addressDecoder : JD.Decoder Address
addressDecoder =
  JD.object4 Address
    (stringField "line1")
    (stringField "line2")
    (stringField "city")
    (stringField "country")


addressEncoder : Address -> JE.Value
addressEncoder v =
  JE.object
    [ ("line1", JE.string v.line1)
    , ("line2", JE.string v.line2)
    , ("city", JE.string v.city)
    , ("country", JE.string v.country)
    ]


type alias User =
  { name : String
  , address : Maybe Address
  }


userDecoder : JD.Decoder User
userDecoder =
  JD.object2 User
    (stringField "name")
    (messageField addressDecoder "address")


userEncoder : User -> JE.Value
userEncoder v =
  JE.object
    [ ("name", JE.string v.name)
    ]


