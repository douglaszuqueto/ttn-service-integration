-- FUNCTION: public.notify_event()

-- DROP FUNCTION public.notify_event();

CREATE FUNCTION public.notify_event()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF 
AS $BODY$

    DECLARE 
        data json;
        notification json;
    
    BEGIN
    
        -- Convert the old or new row to JSON, based on the kind of action.
        -- Action = DELETE?             -> OLD row
        -- Action = INSERT or UPDATE?   -> NEW row
         data = row_to_json(NEW);
       
        
        -- Contruct the notification as a JSON string.
        notification = json_build_object(
                          'table',TG_TABLE_NAME,
                          'action', TG_OP,
                          'data', data);
        
                        
        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('events',notification::text);
        
        -- Result is ignored since this is an AFTER trigger
        RETURN NULL; 
    END;
    
$BODY$;

-- Table: public.metric

-- DROP TABLE public.metric;

CREATE TABLE public.metric
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    app_id character varying COLLATE pg_catalog."default" NOT NULL,
    dev_id character varying COLLATE pg_catalog."default" NOT NULL,
    payload jsonb NOT NULL,
    "time" timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT metric_pkey PRIMARY KEY (id)
)

-- Trigger: products_notify_event

-- DROP TRIGGER products_notify_event ON public.metric;

CREATE TRIGGER products_notify_event
    AFTER INSERT OR UPDATE 
    ON public.metric
    FOR EACH ROW
    EXECUTE PROCEDURE public.notify_event();