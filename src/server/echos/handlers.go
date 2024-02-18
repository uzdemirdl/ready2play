package echos

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/uzdemirdl/ready2play/src/repository"
	"net/http"
)

func getRoot(_ repository.Repository) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "ready2play, podpivas project")
	}
}

func getMap(repos repository.Repository) func(echo.Context) error {
	return func(c echo.Context) error {
		var mapUUID uuid.UUID
		uuidParam := c.Param("uuid")
		if val, err := uuid.Parse(uuidParam); err == nil {
			mapUUID = val
		} else {
			c.Echo().Logger.Error(err)
			return c.String(http.StatusBadRequest, err.Error())
		}
		conn := repos.Conn(context.Background())
		mapValue, err := repos.Maps().GetByUUID(conn, mapUUID)
		if err != nil {
			c.Echo().Logger.Error(err)
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, mapValue)
	}
}

func getMapPoints(repos repository.Repository) func(echo.Context) error {
	return func(c echo.Context) error {
		var mapUUID uuid.UUID
		uuidParam := c.Param("map_uuid")
		if val, err := uuid.Parse(uuidParam); err == nil {
			mapUUID = val
		} else {
			c.Echo().Logger.Error(err)
			return c.String(http.StatusBadRequest, err.Error())
		}
		conn := repos.Conn(context.Background())
		mapPoints, err := repos.MapPoints().GetByMapUUID(conn, mapUUID)
		if err != nil {
			c.Echo().Logger.Error(err)
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, mapPoints)
	}
}
