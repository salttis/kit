# kitti

Työkalujen yhteinen Go-kirjasto (`github.com/salttis/kitti`). Sisältää vain sen, mikä on jaettu vähintään kahden työkalun kesken tai jota seuraava työkalu tarvitsee heti.

| Paketti | Sisältö |
| --- | --- |
| `buildinfo` | Versio, commit, build-aika ja sopimusversio (ldflags) |
| `output` | JSON-kirjekuori `{"schemaVersion":1,"data":...}` ja virhe |
| `contextfile` | Komentajan kontekstitiedosto (`KOMENTAJA_CONTEXT_FILE`) |
| `config` | Asetus- ja välimuistipolut, tiukka TOML-luku |
| `logx` | `log/slog`-lokittaja (stderr, warn; `--verbose` = debug) |
| `ui` | Yhteinen Lip Gloss -teema |

## Säännöt

- Uusi paketti lisätään vasta, kun kaksi työkalua tarvitsee samaa vakaata toimintoa. Ei `utils`-, `common`- tai `helpers`-pakettia.
- Paketit eivät avaa tiedostoja, lue ympäristöä tai ota verkkoyhteyttä `init`-vaiheessa.
- Julkaisu on taggattu semver (`v0.x.y`). Rikkova muutos nostaa version ja kirjataan CHANGELOGiin. Ennen v1.0.0 minor-versio saa rikkoa.
- Työkalut käyttävät `require github.com/salttis/kitti vX.Y.Z` -riviä. `replace`-rivejä ei commitoida; paikallinen kehitys käyttää `~/Tools/go.work`-tiedostoa.
- Charm-versiot (Bubble Tea, Lip Gloss) lukitaan täällä ja työkalut seuraavat niitä.
- Työkalut eivät tuo toistensa pakettia: yhteys on prosessiraja (`--output json`).

## Tarkistus

```powershell
gofmt -l .
go vet ./...
go build ./...
```

## Nimenvaihto 26.11.2026

Kitti v0.3.0 käyttää moduulipolkua `github.com/salttis/kitti`. Työkalut voivat käyttää taggattua versiota ilman paikallista go.work-tiedostoa. Paikalliseen kehitykseen työtila syntyy komennolla `scripts/tools.ps1 sync`. `KOMENTAJA_CONTEXT_FILE` ja `KOMENTAJA_CONTRACT_VERSION` ovat ensisijaiset muuttujat; vanha JAKARI-pari toimii varalla, jos uutta kontekstipolkua ei ole asetettu. Tiedostomuoto ja sopimusversio pysyvät arvossa 1.
