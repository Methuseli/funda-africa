defmodule FundaAfricaWeb.AIChannel do
  use Phoenix.Channel

  def join("ai:tutor:" <> user_id, _params, socket) do
    if socket.assigns.user_id == user_id do
      {:ok, socket}
    else
      {:error, %{reason: "unauthorized"}}
    end
  end

  def handle_in("ai:query", %{"prompt" => prompt}, socket) do
    # TODO: Call FastAPI/gRPC AI service here
    Task.start(fn ->
      push(socket, "ai:status", %{status: "AI is typing..."})
      Process.sleep(1000)
      push(socket, "ai:response", %{response: "Simulated AI answer for: #{prompt}"})
    end)
    {:noreply, socket}
  end
end
