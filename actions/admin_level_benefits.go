package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gophercon/gc18/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (LevelBenefit)
// DB Table: Plural (level_benefits)
// Resource: Plural (AdminLevelBenefits)
// Path: Plural (/admin/level_benefits)
// View Template Folder: Plural (/templates/admin/level_benefits/)

// AdminLevelBenefitsResource is the resource for the level_benefit model
type AdminLevelBenefitsResource struct {
	buffalo.Resource
}

// List gets all LevelBenefits. This function is mapped to the path
// GET /level_benefits
func (v AdminLevelBenefitsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	levelBenefits := &models.LevelBenefits{}
	// You can order your list here. Just change
	err := tx.All(levelBenefits)
	// to:
	// err := tx.Order("create_at desc").All(levelBenefits)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make LevelBenefits available inside the html template
	c.Set("levelBenefits", levelBenefits)
	return c.Render(200, r.HTML("admin/level_benefits/index.html"))
}

// Show gets the data for one LevelBenefit. This function is mapped to
// the path GET /admin/level_benefits/{level_benefit_id}
func (v AdminLevelBenefitsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty LevelBenefit
	levelBenefit := &models.LevelBenefit{}
	// To find the LevelBenefit the parameter admin_level_benefit_id is used.
	err := tx.Find(levelBenefit, c.Param("admin_level_benefit_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make levelBenefit available inside the html template
	c.Set("levelBenefit", levelBenefit)
	return c.Render(200, r.HTML("admin/level_benefits/show.html"))
}

// New renders the formular for creating a new LevelBenefit.
// This function is mapped to the path GET /admin/level_benefits/new
func (v AdminLevelBenefitsResource) New(c buffalo.Context) error {
	// Make levelBenefit available inside the html template
	c.Set("levelBenefit", &models.LevelBenefit{})
	return c.Render(200, r.HTML("admin/level_benefits/new.html"))
}

// Create adds a LevelBenefit to the DB. This function is mapped to the
// path POST /admin/level_benefits
func (v AdminLevelBenefitsResource) Create(c buffalo.Context) error {
	// Allocate an empty LevelBenefit
	levelBenefit := &models.LevelBenefit{}
	// Bind levelBenefit to the html form elements
	err := c.Bind(levelBenefit)
	if err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(levelBenefit)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make levelBenefit available inside the html template
		c.Set("levelBenefit", levelBenefit)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("admin/level_benefits/new.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "LevelBenefit was created successfully")
	// and redirect to the level_benefits index page
	return c.Redirect(302, "/admin/level_benefits/%s", levelBenefit.ID)
}

// Edit renders a edit formular for a level_benefit. This function is
// mapped to the path GET /admin/level_benefits/{level_benefit_id}/edit
func (v AdminLevelBenefitsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty LevelBenefit
	levelBenefit := &models.LevelBenefit{}
	err := tx.Find(levelBenefit, c.Param("admin_level_benefit_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make levelBenefit available inside the html template
	c.Set("levelBenefit", levelBenefit)
	return c.Render(200, r.HTML("admin/level_benefits/edit.html"))
}

// Update changes a level_benefit in the DB. This function is mapped to
// the path PUT /level_benefits/{level_benefit_id}
func (v AdminLevelBenefitsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty LevelBenefit
	levelBenefit := &models.LevelBenefit{}
	err := tx.Find(levelBenefit, c.Param("admin_level_benefit_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Bind LevelBenefit to the html form elements
	err = c.Bind(levelBenefit)
	if err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(levelBenefit)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make levelBenefit available inside the html template
		c.Set("levelBenefit", levelBenefit)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("admin/level_benefits/edit.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "LevelBenefit was updated successfully")
	// and redirect to the level_benefits index page
	return c.Redirect(302, "/admin/level_benefits/%s", levelBenefit.ID)
}

// Destroy deletes a level_benefit from the DB. This function is mapped
// to the path DELETE /admin/level_benefits/{level_benefit_id}
func (v AdminLevelBenefitsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty LevelBenefit
	levelBenefit := &models.LevelBenefit{}
	// To find the LevelBenefit the parameter level_benefit_id is used.
	err := tx.Find(levelBenefit, c.Param("admin_level_benefit_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	err = tx.Destroy(levelBenefit)
	if err != nil {
		return errors.WithStack(err)
	}
	// If there are no errors set a flash message
	c.Flash().Add("success", "LevelBenefit was destroyed successfully")
	// Redirect to the level_benefits index page
	return c.Redirect(302, "/admin/level_benefits")
}
