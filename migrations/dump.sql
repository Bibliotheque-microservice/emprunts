CREATE TABLE emprunts (
    id_emprunt SERIAL PRIMARY KEY,       -- Identifiant unique pour l'emprunt
    utilisateur_id INT NOT NULL,         -- ID de l'utilisateur (lié au Service des Utilisateurs)
    livre_id INT NOT NULL,               -- ID du livre (lié au Service des Livres)
    date_emprunt DATE NOT NULL,          -- Date de début de l'emprunt
    date_retour_prevu DATE NOT NULL,     -- Date prévue pour le retour
    date_retour_effectif DATE,           -- Date réelle du retour (peut être NULL tant que non rendu)
    date_creation TIMESTAMP DEFAULT NOW(), -- Horodatage de la création de l'emprunt
    UNIQUE(utilisateur_id, livre_id, date_retour_effectif) -- Un utilisateur ne peut pas emprunter le même livre deux fois sans le rendre
);

CREATE TABLE penalites (
    id_penalite SERIAL PRIMARY KEY,       -- Identifiant unique pour la pénalité
    utilisateur_id INT NOT NULL,          -- ID de l'utilisateur (lié au Service des Utilisateurs)
    emprunt_id INT NOT NULL,              -- ID de l'emprunt (lié à la table `emprunts`)
    montant DECIMAL(10, 2) NOT NULL,      -- Montant de la pénalité
    paye BOOLEAN DEFAULT FALSE,           -- Indique si la pénalité a été payée (false par défaut)
    date_calcul TIMESTAMP DEFAULT NOW(),  -- Date à laquelle la pénalité a été calculée
    date_paiement TIMESTAMP,              -- Date à laquelle la pénalité a été payée (peut être NULL si non payée)
    FOREIGN KEY (emprunt_id) REFERENCES emprunts (id_emprunt) ON DELETE CASCADE
);


INSERT INTO emprunts (utilisateur_id, livre_id, date_emprunt, date_retour_prevu, date_retour_effectif)
VALUES
(1, 101, '2024-11-01', '2024-11-15', NULL), -- Emprunt en cours
(2, 102, '2024-10-20', '2024-11-03', '2024-11-02'), -- Emprunt rendu à temps
(3, 103, '2024-10-15', '2024-10-29', '2024-11-01'), -- Emprunt rendu en retard
(4, 104, '2024-10-01', '2024-10-15', NULL), -- Emprunt en cours (en retard)
(1, 105, '2024-11-10', '2024-11-25', NULL); -- Nouvel emprunt en cours


INSERT INTO penalites (utilisateur_id, emprunt_id, montant, paye, date_calcul, date_paiement)
VALUES
(3, 3, 5.00, TRUE, '2024-11-02', '2024-11-05'), -- Pénalité payée
(4, 4, 10.00, FALSE, '2024-11-01', NULL); -- Pénalité non payée pour un retard en cours