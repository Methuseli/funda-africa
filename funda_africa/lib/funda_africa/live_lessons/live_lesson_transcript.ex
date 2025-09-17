defmodule FundaAfrica.LiveLessons.LessonTranscript do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "lesson_transcripts" do
    belongs_to(:lesson, FundaAfrica.LiveLessons.LiveLesson, type: :binary_id)
    field(:user, :binary_id)
    field(:content, :string)
    field(:started_at, :utc_datetime)
    # optional, if chunk spans multiple seconds
    field(:ended_at, :utc_datetime)

    timestamps()
  end 

  def changeset(transcript, attrs) do
    transcript
    |> cast(attrs, [:lesson_id, :user, :content])
    |> validate_required([:lesson_id, :content])
  end
end
