package recebimento
import (
	database "lapasta/database"
   "lapasta/internal/models"
)

type RecebimentoRepository interface {
   CriarRecebimento(recebimento *models.Recebimento) error
   ListarRecebimentos() ([]models.Recebimento, error)
}

type recebimentoRepository struct {
   db *database.SQLStr
}

func NovoRecebimentoRepository(db *database.SQLStr) RecebimentoRepository {
   return &recebimentoRepository{
	   db: db,
   }
}

func (r *recebimentoRepository) CriarRecebimento(recebimento *models.Recebimento) error {
   // Reutilizando a implementação existente
   return r.db.CriarRecebimento(recebimento)
}

func (r *recebimentoRepository) ListarRecebimentos() ([]models.Recebimento, error) {
   // Reutilizando a implementação existente
   return r.db.ListarRecebimentos()
}  