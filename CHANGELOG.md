# Muutosloki

## 0.3.0 (26.11.2026)

- Kirjaston nimi vaihtui veska → kitti; uusi moduulipolku on github.com/salttis/kitti. Aiemmat tagit säilyvät vanhalla moduulipolulla.
- contextfile lukee ensisijaisesti KOMENTAJA_CONTEXT_FILE- ja KOMENTAJA_CONTRACT_VERSION-muuttujat. Vanha JAKARI-pari toimii varalla, kun uutta kontekstipolkua ei ole asetettu. Sopimusversio säilyy arvossa 1.

## 0.2.1 (22.11.2026)

- **Rikkova (korjaa 0.2.0:n):** sateenvarjon nimeksi tuli `jakari`, joten `contextfile`-paketin ympäristömuuttujat ovat `JAKARI_CONTEXT_FILE` ja `JAKARI_CONTRACT_VERSION`. Versio 0.2.0 (`SONTSA_*`) oli julkaistuna vain muutaman minuutin eikä ollut kenenkään käytössä; älä käytä sitä.

## 0.2.0 (22.11.2026)

- **Rikkova:** kirjasto nimettiin uudelleen `kit` → `veska`; moduulipolku on `github.com/salttis/veska`.
- Ympäristömuuttujat nimettiin `SONTSA_*` (korvattu 0.2.1:ssä).

## 0.1.0

- Ensimmäinen runko: buildinfo, output, contextfile, config, logx, ui (moduulina `github.com/salttis/kit`; ympäristömuuttujat `JS_*`).
