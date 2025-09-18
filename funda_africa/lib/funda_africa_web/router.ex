defmodule FundaAfricaWeb.Router do
  use FundaAfricaWeb, :router

  pipeline :browser do
    plug(:accepts, ["html"])
    plug(:fetch_session)
    plug(:fetch_live_flash)
    plug(:put_root_layout, html: {FundaAfricaWeb.Layouts, :root})
    plug(:protect_from_forgery)
    plug(:put_secure_browser_headers)
  end

  pipeline :api do
    plug(:accepts, ["json"])
    plug PhoenixSwagger.Plug.Validate
  end

  pipeline :auth do
    plug(FundaAfricaWeb.Plugs.Auth)
  end

  scope "/api/swagger" do
    forward("/", PhoenixSwagger.Plug.SwaggerUI,
      otp_app: :funda_africa,
      swagger_file: "swagger.json"
    )
  end

  scope "/api/v1", FundaAfricaWeb do
    pipe_through([:api, :auth])

    resources("/conversations", ConversationController, only: [:index, :show, :create, :delete])
    resources("/live-lessons", LiveLessonController, only: [:create, :show, :update])
    get("/live-lessons/host", LiveLessonController, :list_by_host)
    get("/live-lessons/participant", LiveLessonController, :list_by_participant)
    get("/conversations/ai", AIInteractionController, :index)
    get("/conversations/ai/:conversation_id", AIInteractionController, :show )
    get("/conversations/:id/messages", MessageController, :index)
  end

  # Enable LiveDashboard and Swoosh mailbox preview in development
  if Application.compile_env(:funda_africa, :dev_routes) do
    # If you want to use the LiveDashboard in production, you should put
    # it behind authentication and allow only admins to access it.
    # If your application does not have an admins-only section yet,
    # you can use Plug.BasicAuth to set up some basic authentication
    # as long as you are also using SSL (which you should anyway).
    import Phoenix.LiveDashboard.Router

    scope "/dev" do
      pipe_through(:browser)

      live_dashboard("/dashboard", metrics: FundaAfricaWeb.Telemetry)
      forward("/mailbox", Plug.Swoosh.MailboxPreview)
    end
  end
end
