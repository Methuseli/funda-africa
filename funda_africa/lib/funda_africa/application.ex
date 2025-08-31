defmodule FundaAfrica.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      FundaAfricaWeb.Telemetry,
      FundaAfrica.Repo,
      {DNSCluster, query: Application.get_env(:funda_africa, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: FundaAfrica.PubSub},
      # Start the Finch HTTP client for sending emails
      {Finch, name: FundaAfrica.Finch},
      # Start a worker by calling: FundaAfrica.Worker.start_link(arg)
      # {FundaAfrica.Worker, arg},
      # Start to serve requests, typically the last entry
      FundaAfricaWeb.Endpoint
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: FundaAfrica.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    FundaAfricaWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
