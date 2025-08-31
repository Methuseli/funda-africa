defmodule FundaAfrica.Repo do
  use Ecto.Repo,
    otp_app: :funda_africa,
    adapter: Ecto.Adapters.Postgres
end
