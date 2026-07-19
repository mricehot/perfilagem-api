CREATE TABLE furos (
  id uuid PRIMARY KEY,
  leque_id uuid NOT NULL REFERENCES leques(id) ON DELETE CASCADE,
  numero text NOT NULL,
  metragem_esperada numeric NOT NULL,
  metragem_real numeric NOT NULL,
  situacao text NOT NULL DEFAULT 'livre' CHECK (situacao IN ('livre', 'obstruido', 'varado')),
  criado_em timestamptz NOT NULL DEFAULT now(),
  UNIQUE (leque_id, numero)
);