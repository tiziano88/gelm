import Debug
import Effects exposing (Effects)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, on, targetValue)
import Http
import Json.Decode as Decode exposing ((:=))
import StartApp
import Task exposing (..)


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
  text "Hello World"


type Action
  = Nop


update : Action -> Model -> (Model, Effects Action)
update action model =
  (model, Effects.none)
