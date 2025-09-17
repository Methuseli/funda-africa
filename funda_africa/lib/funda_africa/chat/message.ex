defmodule FundaAfrica.Chat.Message do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "messages" do
    belongs_to(:conversation, FundaAfrica.Chat.Conversation, type: :binary_id)
    field(:sender, :binary_id)

    # For text or base64/URL to file
    field(:content, :string)
    # "text", "audio", "file", "video"
    field(:content_type, :string, default: "text")
    field(:status, :string, default: "pending")
    field(:recipients, {:array, :binary_id})

    # Optional metadata for files/videos
    field(:file_name, :string)
    field(:file_size, :integer)
    field(:file_url, :string)

    timestamps()
  end

  def changeset(message, attrs) do
    message
    |> cast(attrs, [
      :conversation_id,
      :sender,
      :content,
      :content_type,
      :file_name,
      :file_size,
      :file_url,
      :recipients
    ])
    |> validate_required([:conversation_id, :sender, :content, :content_type])
    |> validate_inclusion(:content_type, ["text", "audio", "file", "video"])
    |> validate_file_size()
  end

  defp validate_file_size(changeset) do
    case get_field(changeset, :file_size) do
      # 20 MB limit
      size when not is_nil(size) and size > 20_000_000 ->
        add_error(changeset, :file_size, "File too large (max 20 MB)")

      _ ->
        changeset
    end
  end
end
