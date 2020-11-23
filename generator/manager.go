package generator

import "github.com/bonbot-team/bonapi/generator/generators"

type Manager struct {
    Generators map[string]Generator
}

func (mgr *Manager) Init() *Manager {
    mgr.RegisterGenerator(&generators.BonToutou{})
    
    return mgr
}

func (mgr *Manager) Get(name string) *Generator {
    gen, found := mgr.Generators[name]
    
    if !found {
        return nil
    }
    
    return &gen
}

func (mgr *Manager) RegisterGenerator(gen Generator) {
    mgr.Generators[gen.GetName()] = gen
}

func GetMgr() *Manager {
    mgr := &Manager{}
    genMap := make(map[string]Generator)
    
    mgr.Generators = genMap
    
    mgr.Init()
    
    return mgr
}
