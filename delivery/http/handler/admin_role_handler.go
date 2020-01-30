package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/user"
	"io"
	"net/http"
	"strconv"
)

// AdminRoleHandler handles role related http requests
type AdminRoleHandler struct {
	roleSvc user.RoleService
}

// NewAdminRoleHandler creates new AdminRoleHandler
func NewAdminRoleHandler(rs user.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{roleSvc: rs}
}

// GetRoles handles Get /v1/admin/roles requests
func (arh *AdminRoleHandler) GetRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	roles, errs := arh.roleSvc.Roles()
	w.Header().Set("Content-Type", "application/json")
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(roles, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Write(output)

}

// GetRole handles Get /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) GetRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	role, errs := arh.roleSvc.Role(uint(id))
	if len(errs) > 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	output, err := json.MarshalIndent(role, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	w.Write(output)
	return
}

// PostRole handles Post /v1/admin/roles requests
func (arh *AdminRoleHandler) PostRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	w.Header().Set("Content-Type", "application/json")
	body := make([]byte, l)
	_, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Print(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	role := entity.Role{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		fmt.Print(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	srole, err2 := arh.roleSvc.StoreRole(&role)
	if len(err2) > 0 {
		fmt.Print(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	p := fmt.Sprintf("/v1/admin/roles/%d", srole.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return

}

// PutRole handles Put /v1/admin/roles:id requests
func (arh *AdminRoleHandler) PutRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	role, err1 := arh.roleSvc.Role(uint(id))
	if len(err1) > 0 {
		fmt.Print(err1)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	_, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Print(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err2 := json.Unmarshal(body, role)
	if err2 != nil {
		fmt.Print("entered err block 3\n")
		http.Error(w, http.StatusText(404), 404)
		return
	}
	role, err3 := arh.roleSvc.UpdateRole(role)
	if len(err3) > 0 {
		fmt.Print("entered err block 4\n")
		http.Error(w, http.StatusText(404), 404)
		return
	}
	resp, err := json.Marshal(role)
	if err != nil {
		fmt.Print("entered err block 5\n")
		http.Error(w, http.StatusText(404), 404)
		return
	}
	w.Write(resp)
	return

}

// DeleteRole handles /v1/admin/roles:id requests
func (arh *AdminRoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
	_, err2 := arh.roleSvc.DeleteRole(uint(id))
	if len(err2) > 0 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
