defmodule FundaAfrica.Model do
  alias FundaAfrica.Chat.{
    ConversationOps,
    ConversationParticipantOps,
    MessageOps,
    MessageReceiptOps
  }
  alias FundaAfrica.AI.AIInteractionOps
  alias FundaAfrica.LiveLessons.{
    LessonParticipantOps,
    LiveLessonOps
    }

  # Conversations
  defdelegate create_conversation(attrs, participants), to: ConversationOps
  defdelegate add_participant(conv_id, user_id), to: ConversationOps
  defdelegate list_by_user(user_id, page \\ 1, page_size \\ 20), to: ConversationOps

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
  defdelegate list_by_host(host_id), to: LiveLessonOps, as: :list_by_host

  # Lesson Participants
  defdelegate list_lesson_participants, to: LessonParticipantOps
  defdelegate get_lesson_participant!(id), to: LessonParticipantOps
  defdelegate create_lesson_participant(attrs), to: LessonParticipantOps
  defdelegate update_lesson_participant(participant, attrs), to: LessonParticipantOps
  defdelegate delete_lesson_participant(participant), to: LessonParticipantOps
  defdelegate list_by_lesson_participants(lesson_id), to: LessonParticipantOps, as: :list_by_lesson

  # Message Receipts
  defdelegate create_message_receipt(attrs), to: MessageReceiptOps
  defdelegate update_status(message_id, recipient_id, status), to: MessageReceiptOps
  defdelegate list_for_message(message_id), to: MessageReceiptOps, as: :list_for_message
  defdelegate all_received?(message_id), to: MessageReceiptOps, as: :all_received?
end
