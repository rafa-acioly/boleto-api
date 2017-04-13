package parser

import "github.com/beevik/etree"

//Rule é o tipo de dados que faz o de-para de uma query na mensagem para uma key no dicionário de valores
type Rule struct {
	XMLQuery string
	MapKey   string
}

//TranslatorMap contem as regras de tradução e extração de uma mensagem
type TranslatorMap struct {
	rules []Rule
}

//AddRule adiciona uma regra nova no objeto de tradução
func (t *TranslatorMap) AddRule(r Rule) {
	t.rules = append(t.rules, r)
}

//GetRules retorna todas as regras do tradutor
func (t *TranslatorMap) GetRules() []Rule {
	return t.rules
}

//NewTranslatorMap cria uma novo objeto de tradução
func NewTranslatorMap() *TranslatorMap {
	t := TranslatorMap{}
	t.rules = make([]Rule, 0, 0)
	return &t
}

//ParseXML executa o parse do XML e retorna uma estrutura de arvore do documento
func ParseXML(xmlDoc string) (*etree.Document, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlDoc); err != nil {
		return etree.NewDocument(), err
	}
	return doc, nil
}

//ExtractValuesFromXML extrai os valores do documento de acordo com uma lista de regras de "de-para"
func ExtractValuesFromXML(doc *etree.Document, translate *TranslatorMap) map[string]string {
	values := make(map[string]string)
	for _, rule := range translate.rules {
		for _, t := range doc.FindElements(rule.XMLQuery) {
			values[rule.MapKey] = t.Text()
			break
		}
	}
	return values
}

// ExtractValues extrai valores de uma string
func ExtractValues(xmlDoc string, translator *TranslatorMap) (map[string]string, error) {
	doc, err := ParseXML(xmlDoc)
	if err != nil {
		return nil, err
	}
	return ExtractValuesFromXML(doc, translator), nil
}
