CREATE TABLE leques (
  id uuid PRIMARY KEY,
  anel_id uuid NOT NULL REFERENCES aneis(id) ON DELETE CASCADE,
  tipo text NOT NULL,
  numero text NOT NULL,
  nome text,
  status text NOT NULL DEFAULT 'aberto',
  criado_em timestamptz NOT NULL DEFAULT now()
);