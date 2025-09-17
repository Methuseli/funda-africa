defmodule FundaAfricaWeb.Presence do
  use Phoenix.Presence,
    otp_app: :funda_africa,
    pubsub_server: FundaAfrica.PubSub
end
