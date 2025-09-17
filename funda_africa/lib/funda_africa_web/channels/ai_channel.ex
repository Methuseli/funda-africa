defmodule FundaAfricaWeb.AIChannel do
  use Phoenix.Channel
  alias FundaAfrica.AI.AIInteractionOps

  # Only the correct user can join their AI channel
  def join("ai:tutor:" <> user_id, _params, socket) do
    if socket.assigns.user_id == user_id do
      {:ok, socket}
    else
      {:error, %{reason: "unauthorized"}}
    end
  end

  # Handle user prompt
  def handle_in("ai:query", %{"prompt" => prompt}, socket) do
    user_id = socket.assigns.user_id

    # Persist the user prompt immediately with empty response
    {:ok, interaction} =
      AIInteractionOps.create_ai_interaction(%{
        user: user_id,
        prompt: prompt,
        response: "",
        model_used: nil,
        metadata: %{}
      })

    # Start async task for AI response
    Task.start(fn ->
      # Notify user that AI is processing
      push(socket, "ai:status", %{status: "AI is typing..."})

      # TODO: replace with real call to FastAPI/gRPC AI service
      Process.sleep(1000)
      ai_response = "Simulated AI answer for: #{prompt}"
      model_used = "simulated-model-v1"

      # Save AI response to the database
      AIInteractionOps.update_ai_interaction(interaction, %{
        response: ai_response,
        model_used: model_used
      })

      # Push the response back to the client
      push(socket, "ai:response", %{
        response: ai_response,
        model_used: model_used,
        interaction_id: interaction.id
      })
    end)

    {:noreply, socket}
  end
end
  