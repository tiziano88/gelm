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
  , s : Maybe Gelm.Message
  }


newModel : Model
newModel =
  { test = 123
  , s = Nothing
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
    ]


type Action
  = Nop
  | Load
  | Resp (Maybe Gelm.Message)


update : Action -> Model -> (Model, Effects Action)
update action model =
  case (Debug.watch "action" action) of
    Nop ->
      (model, Effects.none)

    Resp x ->
      ({ model | s = x }, Effects.none)

    Load ->
      ({ model | test = 222 }, get)

get : Effects Action
get =
  Http.get Gelm.messageDecoder "/api/"
    |> Task.toMaybe
    |> Task.map Resp
    |> Effects.task
