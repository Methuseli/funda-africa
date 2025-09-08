defmodule FundaAfrica.Chat do
  alias FundaAfrica.Chat.{
    ConversationOps,
    ConversationParticipantOps,
    MessageOps,
    AIInteractionOps,
    LiveLessonOps,
    LessonParticipantOps
  }

  # Conversations
  defdelegate list_conversations, to: ConversationOps
  defdelegate get_conversation!(id), to: ConversationOps
  defdelegate create_conversation(attrs), to: ConversationOps
  defdelegate update_conversation(conversation, attrs), to: ConversationOps
  defdelegate delete_conversation(conversation), to: ConversationOps
  defdelegate list_by_creator(user_id), to: ConversationOps

  # Conversation Participants
  defdelegate list_conversation_participants, to: ConversationParticipantOps
  defdelegate get_conversation_participant!(id), to: ConversationParticipantOps
  defdelegate create_conversation_participant(attrs), to: ConversationParticipantOps
  defdelegate update_conversation_participant(participant, attrs), to: ConversationParticipantOps
  defdelegate delete_conversation_participant(participant), to: ConversationParticipantOps
  defdelegate list_by_conversation(conversation_id), to: ConversationParticipantOps

  # Messages
  defdelegate list_messages, to: MessageOps
  defdelegate get_message!(id), to: MessageOps
  defdelegate create_message(attrs), to: MessageOps
  defdelegate update_message(message, attrs), to: MessageOps
  defdelegate delete_message(message), to: MessageOps
  defdelegate list_by_conversation_messages(conversation_id), to: MessageOps, as: :list_by_conversation

  # AI Interactions
  defdelegate list_ai_interactions, to: AIInteractionOps
  defdelegate get_ai_interaction!(id), to: AIInteractionOps
  defdelegate create_ai_interaction(attrs), to: AIInteractionOps
  defdelegate update_ai_interaction(ai_interaction, attrs), to: AIInteractionOps
  defdelegate delete_ai_interaction(ai_interaction), to: AIInteractionOps
  defdelegate list_by_user_ai(user_id), to: AIInteractionOps, as: :list_by_user

  # Live Lessons
  defdelegate list_live_lessons, to: LiveLessonOps
  defdelegate get_live_lesson!(id), to: LiveLessonOps
  defdelegate create_live_lesson(attrs), to: LiveLessonOps
  defdelegate update_live_lesson(lesson, attrs), to: LiveLessonOps
  defdelegate delete_live_lesson(lesson), to: LiveLessonOps
  defdelegate list_by_teacher_lessons(teacher_id), to: LiveLessonOps, as: :list_by_teacher

  # Lesson Participants
  defdelegate list_lesson_participants, to: LessonParticipantOps
  defdelegate get_lesson_participant!(id), to: LessonParticipantOps
  defdelegate create_lesson_participant(attrs), to: LessonParticipantOps
  defdelegate update_lesson_participant(participant, attrs), to: LessonParticipantOps
  defdelegate delete_lesson_participant(participant), to: LessonParticipantOps
  defdelegate list_by_lesson_participants(lesson_id), to: LessonParticipantOps, as: :list_by_lesson
end
