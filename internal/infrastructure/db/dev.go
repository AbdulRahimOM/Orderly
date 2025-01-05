package db

import (
	"flag"
	"fmt"
	"orderly/internal/domain/response"
	"orderly/internal/infrastructure/config"

	"github.com/gofiber/fiber/v2"
)

func resetDBFlagCalled() bool {
	init_db_condition := flag.Bool("resetdb", false, "to ask if the database should be initialized empty")
	flag.Parse()
	if *init_db_condition {
		fmt.Println("'resetdb' flag has been used")
		return true
	}
	return false
}

// ResetDB is a handler to reset the database. Its for development purpose only.
func ResetDB(c *fiber.Ctx) error {

	err := ClearDB()
	if err != nil {
		return response.ErrorResponse(500, "failed to clear the database:", err).WriteToJSON(c)
	}

	InitDB()

	fmt.Println("Database reset successful")
	return response.SuccessResponse(200, "Database reset successfully", nil).WriteToJSON(c)
}

func ClearDB() error {
	var dropSQL string
	switch config.Configs.Env.Environment {
	case "LOCAL", "DOCKER":
		// Drop all tables and objects
		dropSQL = `
    DO $$ 
    DECLARE
        r RECORD;
        s RECORD;
    BEGIN
        -- Loop over all schemas except system schemas
        FOR s IN (SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('pg_catalog', 'information_schema', 'pg_temp_X', 'pg_toast', 'pg_audit')) LOOP
            
            -- Disable referential integrity checks temporarily for this schema
            EXECUTE 'SET session_replication_role = replica';
            -- Drop all tables in the current schema
            FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = s.schema_name) LOOP
                EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(s.schema_name) || '.' || quote_ident(r.tablename) || ' CASCADE';
            END LOOP;
            -- Drop all sequences in the current schema
            FOR r IN (SELECT sequencename FROM pg_sequences WHERE schemaname = s.schema_name) LOOP
                EXECUTE 'DROP SEQUENCE IF EXISTS ' || quote_ident(s.schema_name) || '.' || quote_ident(r.sequencename) || ' CASCADE';
            END LOOP;
            -- Re-enable referential integrity checks for this schema
            EXECUTE 'SET session_replication_role = DEFAULT';
        END LOOP;
    END $$;
    `
	// case "RENDER":
	default:
		dropSQL = `
	DO $$
	DECLARE
		r RECORD;
		s RECORD;
	BEGIN
		-- Loop over all schemas except system schemas
		FOR s IN (SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('pg_catalog', 'information_schema', 'pg_temp_X', 'pg_toast', 'pg_audit')) LOOP
			
			-- Drop all tables in the current schema
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = s.schema_name) LOOP
				BEGIN
					EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(s.schema_name) || '.' || quote_ident(r.tablename) || ' CASCADE';
				EXCEPTION WHEN others THEN
					RAISE NOTICE 'Error dropping table %: %', r.tablename, SQLERRM;
				END;
			END LOOP;
			-- Drop all sequences in the current schema
			FOR r IN (SELECT sequencename FROM pg_sequences WHERE schemaname = s.schema_name) LOOP
				BEGIN
					EXECUTE 'DROP SEQUENCE IF EXISTS ' || quote_ident(s.schema_name) || '.' || quote_ident(r.sequencename) || ' CASCADE';
				EXCEPTION WHEN others THEN
					RAISE NOTICE 'Error dropping sequence %: %', r.sequencename, SQLERRM;
				END;
			END LOOP;
		END LOOP;
	END $$;
	`
	}

	// Execute the SQL
	err := DB.Exec(dropSQL).Error
	if err != nil {
		// return response.CreateError(500, "failed to clear the database:", err).WriteToJSON(c)
		return fmt.Errorf("failed to clear the public schema data. error: %v", err)
	}

	dropSchemasSQL := `
    DO $$ 
    DECLARE
        s RECORD;
    BEGIN
        -- Loop over all user-defined schemas
        FOR s IN (SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ('pg_catalog', 'information_schema', 'pg_temp_X', 'pg_toast', 'pg_audit', 'public')) LOOP
            EXECUTE 'DROP SCHEMA IF EXISTS ' || quote_ident(s.schema_name) || ' CASCADE';
        END LOOP;
    END $$;
    `

	// Execute the SQL to drop user-defined schemas
	err = DB.Exec(dropSchemasSQL).Error
	if err != nil {
		// return response.CreateError(500, "failed to clear the database:", err).WriteToJSON(c)
		return fmt.Errorf("failed to clear the database: %v", err)
	}

	// db.InitDB()

	fmt.Println("Database clearing successful")
	// return response.CreateSuccess(200, "Database reset successfully", nil).WriteToJSON(c)
	return nil
}
