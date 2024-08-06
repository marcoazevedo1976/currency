type LinkNumeroPagina struct {
	Numero int
	Link   template.HTML
	Atual  bool
}

type Paginador struct {
	Paginas            []LinkNumeroPagina
	TotalRegistros     int
	LinkPaginaAnterior template.HTML
	LinkProximaPagina  template.HTML
	paginaAtual        int
	totalPaginas       int
	offset             int
	finish             int
}

func (p Paginador) TemPaginaAnterior() bool {
	return p.paginaAtual > 1
}

func (p Paginador) TemProximaPagina() bool {
	return p.paginaAtual < p.totalPaginas
}

func NewPaginador(registrosPorPagina int, paginaAtual int, totalRegistros int) Paginador {
	var p Paginador
	paginasVisiveis := 10
	p.paginaAtual = paginaAtual
	p.offset = registrosPorPagina * (paginaAtual - 1)
	p.finish = p.offset + registrosPorPagina
	if p.finish > totalRegistros {
		p.finish = totalRegistros
	}
	p.TotalRegistros = totalRegistros
	p.totalPaginas = int(math.Ceil(float64(p.TotalRegistros) / float64(registrosPorPagina)))
	ultimaPagina := int(math.Ceil(float64(p.paginaAtual)/float64(paginasVisiveis))) * paginasVisiveis
	primeiraPagina := (ultimaPagina - paginasVisiveis) + 1
	if ultimaPagina > p.totalPaginas {
		ultimaPagina = p.totalPaginas
	}
	for idx := primeiraPagina; idx <= ultimaPagina; idx++ {
		atual := idx == paginaAtual
		p.Paginas = append(p.Paginas, LinkNumeroPagina{Numero: idx, Atual: atual})
	}
	return p
}
