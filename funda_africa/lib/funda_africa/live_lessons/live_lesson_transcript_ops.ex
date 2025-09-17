defmodule FundaAfrica.LiveLessons.LessonTranscriptOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.LiveLessons.LessonTranscript

  def create_transcript(attrs) do
    %LessonTranscript{}
    |> LessonTranscript.changeset(attrs)
    |> Repo.insert()
  end

  def list_by_lesson(lesson_id) do
    LessonTranscript
    |> where(lesson_id: ^lesson_id)
    |> Repo.all()
  end
end
