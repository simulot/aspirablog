package blogger

import (
	"testing"
)

type FakeGetter struct {
}

func (g *FakeGetter) Get(url string) ([]byte, string, error) {
	return nil, "", nil
}
func (g *FakeGetter) IsUpdated(url string) (bool, error) {
	return false, nil
}

const text = `<div class="separator" style="clear: both; text-align: center;"></div><div class="separator" style="clear: both; text-align: center;"><a href="https://4.bp.blogspot.com/-WLE8QsGuRvM/W319wgyqVEI/AAAAAAAA90A/xZBFg2lpyEA3saMicVMguBHZb681kv35QCKgBGAs/s1600/20180814_124909.jpg" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" data-original-height="900" data-original-width="1600" height="360" src="https://4.bp.blogspot.com/-WLE8QsGuRvM/W319wgyqVEI/AAAAAAAA90A/xZBFg2lpyEA3saMicVMguBHZb681kv35QCKgBGAs/s640/20180814_124909.jpg" width="640" /></a></div><br /><br /><br /><b>LE&nbsp; MARCHE :</b><br /><b><br /></b><br /><ul><li>600 g de poitrine d'agneau</li><li>2 pommes de terre par personne</li><li>5 à 6 feuilles de sauge</li><li>sel, poivre et huile d'olive</li></ul><br /><b><br /></b><b><br /></b><b>LA&nbsp;&nbsp;MÉTHODE :</b><br /><b><br /></b><b><br /></b>Laver et couper les pommes de terre à la mandoline sans les peler et les installer tout de suite sur la plaque de cuisson.<br /><br /><div class="separator" style="clear: both; text-align: center;"></div><br /><div class="separator" style="clear: both; text-align: center;"><a href="https://4.bp.blogspot.com/-w0ydEV1b_tQ/W3w2N8mjSbI/AAAAAAAA9yk/wiPcJ_W5cmgBSkwDuOT8_Iscrg_HMaFIQCKgBGAs/s1600/20180814_115004.jpg" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" data-original-height="900" data-original-width="1600" height="225" src="https://4.bp.blogspot.com/-w0ydEV1b_tQ/W3w2N8mjSbI/AAAAAAAA9yk/wiPcJ_W5cmgBSkwDuOT8_Iscrg_HMaFIQCKgBGAs/s400/20180814_115004.jpg" width="400" /></a></div><br />Assaisonner et mettre un petit filet d'huile d'olive.<br />Poser les feuilles de sauge.<br /><br /><div class="separator" style="clear: both; text-align: center;"></div><br /><div class="separator" style="clear: both; text-align: center;"><a href="https://3.bp.blogspot.com/-C3bp9A21lLY/W3w2ae3ZtZI/AAAAAAAA9yo/vq4JyTTfA20JZlPoC3tHQ4Lh_tK9kmAWgCKgBGAs/s1600/20180814_115410.jpg" imageanchor="1" style="margin-left: 1em; margin-right: 1em;"><img border="0" data-original-height="900" data-original-width="1600" height="225" src="https://3.bp.blogspot.com/-C3bp9A21lLY/W3w2ae3ZtZI/AAAAAAAA9yo/vq4JyTTfA20JZlPoC3tHQ4Lh_tK9kmAWgCKgBGAs/s400/20180814_115410.jpg" width="400" /></a></div><br />Déposer l'agneau assaisonné et passer au four à chaleur tournante à 180° pendant 25 minutes environ.<br /><br /><b><br /></b><b>MON&nbsp; ASTUCE :</b><br /><b><br /></b>J'interprète cette version très simple, très rapide et très économique de diverses façons :<br />Un peu moins économique mais tout aussi simple et rapide !!<br /><br /><table align="center" cellpadding="0" cellspacing="0" class="tr-caption-container" style="margin-left: auto; margin-right: auto; text-align: center;"><tbody><tr><td style="text-align: center;"><a href="https://1.bp.blogspot.com/-_7wDVfHONHw/W3w4jpADsOI/AAAAAAAA9y0/1oYm8Xn14dge6ObPLRn6wnKlSyA27pSOQCKgBGAs/s1600/20180819_125429.jpg" imageanchor="1" style="margin-left: auto; margin-right: auto;"><img border="0" data-original-height="900" data-original-width="1600" height="225" src="https://1.bp.blogspot.com/-_7wDVfHONHw/W3w4jpADsOI/AAAAAAAA9y0/1oYm8Xn14dge6ObPLRn6wnKlSyA27pSOQCKgBGAs/s400/20180819_125429.jpg" width="400" /></a></td></tr><tr><td class="tr-caption" style="text-align: center;">Ce jour là, c'était avec des tranches de gigot</td></tr></tbody></table>L'autre jour ,c'était avec des côtelettes..<br /><br /><div class="separator" style="clear: both; text-align: center;"><b></b></div>`

func Test_ParseHTML(t *testing.T) {
	b := New("test", "test url", "", &FakeGetter{})

	p, err := b.parsePost(text)

	_, _ = p, err
}
