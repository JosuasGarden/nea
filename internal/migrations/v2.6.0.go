package migrations

import (
	"github.com/jmoiron/sqlx"
	"github.com/knadh/koanf/v2"
	"github.com/knadh/listmonk/models"
	"github.com/knadh/stuffbin"
)

// V2_6_0 performs the DB migrations.
func V2_6_0(db *sqlx.DB, fs stuffbin.FileSystem, ko *koanf.Koanf) error {
	// Add new template_type.
	if _, err := db.Exec(`
        ALTER TYPE template_type ADD VALUE IF NOT EXISTS 'product';
	`); err != nil {
		return err
	}

	// Insert new default product template.
	prodTpl, err := fs.Get("/static/email-templates/default-product.tpl")
	if err != nil {
		return err
	}

	var prodTplID int
	if err := db.Get(
		&prodTplID,
		`
            INSERT INTO templates (name, type, subject, body) 
                VALUES($1, $2, $3, $4)
                RETURNING id;
        `,
		"Default product template",
		models.TemplateTypeProduct,
		"",
		prodTpl.ReadBytes(),
	); err != nil {
		return err
	}

	// Remove uniqueness constraint on templates
	if _, err := db.Exec(`DROP INDEX IF EXISTS templates_is_default_idx`); err != nil {
		return err
	}

	// Set default product template to created
	if _, err := db.Exec(`
        UPDATE templates SET is_default=true WHERE id = $1;
	`, prodTplID); err != nil {
		return err
	}

	// Create new product_template_id column in campaigns
	if _, err := db.Exec(`ALTER TABLE campaigns
        ADD COLUMN IF NOT EXISTS product_template_id INTEGER REFERENCES templates(id) ON DELETE SET DEFAULT DEFAULT 2;
    `); err != nil {
		return err
	}

	return nil
}
