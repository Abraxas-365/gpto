package application

func (a *app) SyncCategories(isNew bool) error {

	if isNew {
		//hacer la carga inicial en magento
		if err := a.magento.CategoryInicialLoad(); err != nil {
			return err
		}
	}
	//traer las categorias , y enviarlas a akeneo
	category, err := a.magento.GetCategories()
	if err != nil {
		return err
	}

	//enviar a akeneo
	if err := a.akeneo.UploadCategories(category); err != nil {
		return err
	}

	return nil
}
