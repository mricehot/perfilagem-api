CREATE TABLE aneis (
  id uuid PRIMARY KEY,
  nome text NOT NULL,
  ativo boolean NOT NULL DEFAULT false,
  criado_em timestamptz NOT NULL DEFAULT now()
);