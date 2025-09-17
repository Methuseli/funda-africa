defmodule FundaAfrica.Chat.Memberships do
  @moduledoc """
  Finds all channels a user belongs to (orgs, lessons, groups, private chats).
  """

  import Ecto.Query
  alias FundaAfrica.Repo
  alias FundaAfrica.Chat.{Conversation, ConversationParticipant}
  alias FundaAfrica.LiveLessons.{LiveLesson, LessonParticipant}

  @doc """
  Returns a list of topic strings a user should join on socket connect.
  """
  def topics_for_user(user_id) do
    conv_topics =
      from(c in Conversation,
        join: p in ConversationParticipant, on: p.conversation_id == c.id,
        where: p.user == ^user_id,
        select: {"chat:#{c.type}:#{c.id}"}
      )
      |> Repo.all()

    lesson_topics =
      from(l in LiveLesson,
        join: p in LessonParticipant, on: p.lesson_id == l.id,
        where: p.user == ^user_id,
        select: {"chat:lesson:#{l.id}"}
      )
      |> Repo.all()

    conv_topics ++ lesson_topics
  end
end
