import Debug
import Effects exposing (Effects)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, on, targetValue)
import Http
import Json.Decode as Decode
import StartApp
import Task exposing (..)

import Gelm


app =
  StartApp.start
    { init = init
    , view = view
    , update = update
    , inputs = []
    }


main = app.html


port tasks : Signal (Task.Task Effects.Never ())
port tasks =
  app.tasks


type alias Model =
  { test : Int
  , data : Maybe Gelm.Message
  }


newModel : Model
newModel =
  { test = 123
  , data = Nothing
  }


init : (Model, Effects Action)
init =
  (newModel, get)


(=>) = (,)


view : Signal.Address Action -> Model -> Html
view address model =
  div []
    [ a
      [ onClick address Load ]
      [ text "Hello World 3" ]
    , text <| (toString model)
    , optional (dataWidget address) model.data
    ]


optional : (a -> Html) -> Maybe a -> Html
optional f d =
  Maybe.map f d |> Maybe.withDefault (text "null")


dataWidget : Signal.Address Action -> Gelm.Message -> Html
dataWidget address data =
  div []
    [ input
      [ value data.fieldWithLongName
      , on "input" targetValue (Signal.message address << Update) ] []
    ]


type Action
  = Nop
  | Load
  | Update String
  | Resp (Maybe Gelm.Message)


update : Action -> Model -> (Model, Effects Action)
update action model =
  case (Debug.watch "action" action) of
    Nop ->
      (model, Effects.none)

    Resp x ->
      ({ model | data = x }, Effects.none)

    Update f ->
      case model.data of
        Just data ->
          ({ model | data = Just { data | fieldWithLongName = f } }, Effects.none)
        Nothing ->
          (model, Effects.none)

    Load ->
      ({ model | test = 222 }, get)

get : Effects Action
get =
  Http.get Gelm.messageDecoder "/api/"
    |> Task.toMaybe
    |> Task.map Resp
    |> Effects.task
