package generator

import "github.com/bonbot-team/bonapi/generator/generators"

type GeneratorManager struct {
    Generators map[string]Generator
}

func (mgr *GeneratorManager) Init() *GeneratorManager {
    mgr.RegisterGenerator(&generators.BonToutou{})
    
    return mgr
}

func (mgr *GeneratorManager) Get(name string) *Generator {
    gen, found := mgr.Generators[name]
    
    if !found {
        return nil
    }
    
    return &gen
}

func (mgr *GeneratorManager) RegisterGenerator(gen Generator) {
    mgr.Generators[gen.GetName()] = gen;
}

func GetMgr() *GeneratorManager {
    mgr := &GeneratorManager{}
    genMap := make(map[string]Generator)
    
    mgr.Generators = genMap
    
    mgr.Init()
    
    return mgr
}
