package dimextract

import (
    "fmt"
    "testing"
)

type TestCase struct {
    desc string
    d1 int
    d2 int
    d3 int
}

func (t *TestCase) str() string {
    return fmt.Sprintf("(%d / %d / %d )", t.d1, t.d2, t.d3)
}

func (t *TestCase) equal(d *Dimension) bool {
    return (t.d1 == d.d1) && (t.d2 == d.d2) && (t.d3 == d.d3)
}

func doTest(t *testing.T, testCases []TestCase) {
    for _, tc := range testCases {
        d, err := ExtractDims(tc.desc)
        if err != nil {
            t.Error(err)
        }
        if !tc.equal(&d) {
            t.Errorf("Missmatch  %s, got : %s on %s", tc.str(), d.str(), tc.desc)
            t.Fail()
        } else {

        }
    }
}

func TestExtractSimple(t *testing.T) {

    testCases := []TestCase{
        {"13 x 12 x 11", 13, 12, 11,},
        {"13 x 12 x 11 cm", 13, 12, 11,},
        {"13 x 12 cm", 13, 12, 0,},
        {"13 cm", 13, 0, 0,},
    }

    doTest(t, testCases)
}

func TestExtractHLLPrefix(t *testing.T) {

    testCases := []TestCase{
        {"obes en verre givré. Haut.: 42 cm", 42, 0, 0},
        {"bakélite et métal. Circa 1960. Dans son étui. Long.: 27 cm", 27, 0, 0},
        {"Style Louis XV. Haut.: 95 cm Long.: 75 Prof.: 42 cm (usure)", 95, 75, 42},
        {"bles. Repose sur des roulettes. Haut.: 63 cm Long.: 60 cm Larg.: 36 cm", 63, 60, 36},
        {"nture en bronze et laiton. Longueur 38 cm", 38, 0, 0},
        {"Vase en opaline blanche, monté en lampe. Haut.: 25 cm", 25, 0, 0},
    }

    doTest(t, testCases)
}

func TestExtract(t *testing.T) {
    testCases := []TestCase{
        {"Blabla bla 13 * 12 * 11 cm", 13, 12, 11,},
        {" , signé dans la planche. Encadré sous verre. 50X 64 cm à vue", 64, 50, 0,},
        {". En feuille,18 x28 cm ( légères rousseurs)", 28, 18, 0,},
        {"XIXe siècle. Dimensions :72 X 68 cm", 72, 68, 0},
        {" Encadré sous verre. 50X 64 cm à vue", 64, 50, 0},
        {" en porcelaine. 8-14 cm", 14, 8, 0},
        {" Accessoire de cordonnier en fonte. 16 cm", 16, 0, 0},
        {"Dans le goût de Murano, Vase en verre coloré. Haut.: 19 cm Circa 1970.", 19, 0, 0},
        {"Lot comprenant une paire de voilages ( 140 x 300)  et deux voilages d''un modèle proche ( 140 x 280 cm). Etat neuf et TBE", 280, 140, 0},
        {"ies - Rayon Bagage\" 60 x 80 x 50 cm Début Xxe. (manque un élément d''une des serrures)", 80, 60, 50},
        {"Malle en bois et métal. 41 x 85 x 51 cm (usure, traces de colle).", 85, 51, 41},
        {"On y joint un plat en faïence blanche. 25-33 cm", 33, 25, 0},
        {"ORIENT Galerie en laine rouge à décor de végétaux,340 X 80 cm", 340, 80, 0},
        {"IRAN Grand tapis en laine fond rouge à motif de végétaux stylisés.207X 287 CM", 287, 207, 0},
        {"nnes en lin, brodées et ajourées, chiffrées (80 x 80 cm)", 80, 80, 0},
        {"Boîte en argent étranger (titre 800), très ouvragée. Poids : 179 g 4 x 10 x 8 cm", 10, 8, 4},
        {"Carré en soie à décor de ceintures sur fond blanc et bordure bleue. TBE. 81 X 81 CM", 81, 81, 0},
        {"Lot de deux colliers en perles fantaisie. Long. (ouverts) : 40-48 cm Dans une pochette en cuir noir.", 48, 40, 0},
        {"Paire de bougeoirs en bronze, modèle Rocaille. Montés en lampe, avec leurs abat-jours. Hauteur totale : 28 cm", 28, 0, 0},
        {"Circa 1960. 19 x 26 x 7 cm TBE.", 26, 19, 7},
        {"Lot de coupons de tissu d'ameublement. De 34 à 110 cm", 110, 0, 0},
        {"à décor de végétaux.Circa 1930.H,21 cm", 21, 0, 0},
        {"Lot de trois assiettes en porcelaine blanche.D.27 cm", 27, 0, 0},
    }

    doTest(t, testCases)
}

func _TestFailing(t *testing.T) {
    testCases := []TestCase{
        {"Début Xxeme siècle .D.30 cm ( égrenures , cassé/collé)", 30, 30, 0,},
        {"(accident).Circa 1930. H.totale 57 cm", 57, 0, 0,},
        {"Environ 7 cm", 7, 0, 0,},
        {"Circa 1960. 19 x 26 x 7 cm TBE.", 26, 19, 7},
        {"Lot de coupons de tissu d'ameublement. De 34 à 110 cm", 110, 34, 0},
        {"à décor de végétaux.Circa 1930.H,21 cm", 21, 0, 0},
        {"Lot de trois assiettes en porcelaine blanche.D.27 cm", 27, 27, 0},
        {"Hauteur assise : 45 cm, hauteur totale : 82 cm (importante usure au vernis)", 82, 0, 0},
        {"PIERROT LE LOUP Manteau long en vison. Largeur épaules : 46 cm Long.: 90 cm", 90, 46, 0},
        {"SKYBOLTE Paire de jumelles en métal laqué noir. 10x40, 105m/1000m. Larg.: 17 cm Dans leur étui (acc).", 40, 10, 0},
        {"Lot de trois assiettes en porcelaine blanche.D.27 cm", 27, 27, 0},
        {"r, bracelet cuir. Diamètre du cadran : 4 cm", 4, 4, 0},
        {"à décor de fleurs : 4 paires différentes (hauteur environ 150 cm). On y joint 2 stores.", 150, 0, 0},
        {"Lot de deux cadres en bois et velours, dont un avec application de laiton. 28 x 22 et 35 x 25 cm", 47, 35, 0},
        {"SARREGUEMINE Lot comprenant un vide-poche en faïence verte en forme de feuille (égrenure) et un moutardier. 17 et 6 cm", 17, 6, 0},
        {"CHAMPAGNE Lot publicitaire comprenant un seau en aluminium \"Moet et Chandon\" (20 cm) et 3 coupelles en faïence d''Orchies \"Eugène Cliquot\" (diam.: 10 cm).", 20, 20, 10},
        {"Série de 2 tables gigognes en laiton et verre. Haut.: 40-43 cm Long.: 55 cm On y joint une table basse en laiton et marbre (43x82x45 cm)", 82, 43, 36},
    }

    doTest(t, testCases)
}