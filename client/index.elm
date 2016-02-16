import Debug
import Effects exposing (Effects)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, on, targetValue)
import Http
import Json.Decode as Decode exposing ((:=))
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


type alias Model =
  { test : Int
  }


newModel : Model
newModel =
  { test = 123
  }


init : (Model, Effects Action)
init =
  (newModel, Effects.none)


(=>) = (,)


view : Signal.Address Action -> Model -> Html
view address model =
  div
    [ onClick address Load ]
    [ text "Hello World 2" ]


type Action
  = Nop
  | Load


update : Action -> Model -> (Model, Effects Action)
update action model =
  case action of
    Nop ->
      (model, Effects.none)

    Load ->
      (model, Http.get Decode.string "/api/" |> Task.toMaybe |> Task.map (\_ -> Nop) |> Effects.task)
