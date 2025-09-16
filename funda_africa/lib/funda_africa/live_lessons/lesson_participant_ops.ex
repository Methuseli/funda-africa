defmodule FundaAfrica.LiveLessons.LessonParticipantOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.LiveLessons.LessonParticipant

  def list_lesson_participants, do: Repo.all(LessonParticipant)
  def get_lesson_participant!(id), do: Repo.get!(LessonParticipant, id)

  def create_lesson_participant(attrs \\ %{}) do
    %LessonParticipant{}
    |> LessonParticipant.changeset(attrs)
    |> Repo.insert()
  end

  def update_lesson_participant(%LessonParticipant{} = participant, attrs) do
    participant
    |> LessonParticipant.changeset(attrs)
    |> Repo.update()
  end

  def delete_lesson_participant(%LessonParticipant{} = participant), do: Repo.delete(participant)

  def list_by_lesson(lesson_id) do
    LessonParticipant
    |> where(lesson_id: ^lesson_id)
    |> Repo.all()
  end
end
