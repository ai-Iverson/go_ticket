package logic

import (
	_ "go_ticket/internal/logic/casbin"
	_ "go_ticket/internal/logic/context"
	_ "go_ticket/internal/logic/cron"
	_ "go_ticket/internal/logic/knowledge"
	_ "go_ticket/internal/logic/middleware"
	_ "go_ticket/internal/logic/ticket"
	_ "go_ticket/internal/logic/user"
)
