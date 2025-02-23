package controllers

import (
	"encoding/json"
	"net/http"
    "github.com/gorilla/mux"

	"github.com/xDani-v/umec_api_goolang/data"
	"github.com/xDani-v/umec_api_goolang/models"
	"github.com/xDani-v/umec_api_goolang/utils"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var roles []models.Rol
	data.DB.Find(&roles)

	respuesta := utils.ResponseMsg{
		Msg:  "Roles",
		Data: roles,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}
func GetRolesPaginated(w http.ResponseWriter, r *http.Request) {
    var roles []models.Rol
    pagination := models.Pagination{
        Page:     utils.StringToInt(r.URL.Query().Get("page"), 1),
        Limit:    utils.StringToInt(r.URL.Query().Get("limit"), 10),
        Search:   r.URL.Query().Get("search"),
        OrderBy:  r.URL.Query().Get("orderBy"),
        OrderDir: r.URL.Query().Get("orderDir"),
    }
    
    // Initialize query without executing it
    query := data.DB

    // Apply search filter
    if pagination.Search != "" {
        query = query.Where("nombre LIKE ?", "%"+pagination.Search+"%")
    }

    // Count total rows with search applied
    query.Model(&models.Rol{}).Count(&pagination.TotalRows)

    // Apply ordering
    if pagination.OrderBy != "" {
        if pagination.OrderDir != "desc" {
            pagination.OrderDir = "asc"
        }
        query = query.Order(pagination.OrderBy + " " + pagination.OrderDir)
    } else {
        query = query.Order("id asc")
    }

    // Apply pagination and execute query
    offset := (pagination.Page - 1) * pagination.Limit
    query.Offset(offset).Limit(pagination.Limit).Find(&roles)

    result := models.PaginationResult{
        Items:      roles,
        Page:       pagination.Page,
        Limit:      pagination.Limit,
        TotalRows:  pagination.TotalRows,
        TotalPages: int(pagination.TotalRows)/pagination.Limit + 1,
    }

    respuesta := utils.ResponseMsg{
        Msg:    "Roles paginados",
        Data:   result,
        Status: 200,
    }
    w.Header().Set(contentType, applicationJSON)
    json.NewEncoder(w).Encode(respuesta)
}

func GetRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	data.DB.First(&rol, r.URL.Query().Get("id"))

	respuesta := utils.ResponseMsg{
		Msg:    "Rol encontrado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func CreateRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	json.NewDecoder(r.Body).Decode(&rol)
	data.DB.Create(&rol)

	respuesta := utils.ResponseMsg{
		Msg:    "Rol creado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func UpdateRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	json.NewDecoder(r.Body).Decode(&rol)
	data.DB.Save(&rol)

	respuesta := utils.ResponseMsg{
		Msg:    "Rol actualizado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func DeleteRol(w http.ResponseWriter, r *http.Request) {
	rol := models.Rol{}
	data.DB.First(&rol, r.URL.Query().Get("id"))
	data.DB.Delete(&rol)

	respuesta := utils.ResponseMsg{
		Msg:    "Rol eliminado",
		Data:   rol,
		Status: 200,
	}
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(respuesta)
}

func RolFuncionalidades(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    rolID := vars["id"]

    var results []map[string]interface{}

    query := `
        SELECT 
            m.id as menu_id,
            m.nombre as menu_nombre,
            m.icono as menu_icon,
            f.id as funcionalidad_id,
            f.nombre as func_nombre,
            f.icono as func_icon,
            f.ruta as func_ruta,
            f.estado as func_estado,
            rf.crear,
            rf.leer,
            rf.actualizar,
            rf.eliminar
        FROM rolesfuncionalidad rf
        INNER JOIN roles r ON r.id = rf.id_rol and r.estado = true
        INNER JOIN funcionalidad f ON f.id = rf.id_funcionalidad and f.estado = true
        INNER JOIN menu m ON m.id = f.id_menu and m.estado = true
        WHERE rf.id_rol = ? order by m.id asc
    `

    if err := data.DB.Raw(query, rolID).Scan(&results).Error; err != nil {
        respuesta := utils.ResponseMsg{
            Msg:    "Error al obtener funcionalidades",
            Status: http.StatusInternalServerError,
        }
        w.Header().Set(contentType, applicationJSON)
        json.NewEncoder(w).Encode(respuesta)
        return
    }

    respuesta := utils.ResponseMsg{
        Msg:    "Funcionalidades del rol",
        Data:   results,
        Status: http.StatusOK,
    }
    w.Header().Set(contentType, applicationJSON)
    json.NewEncoder(w).Encode(respuesta)
}