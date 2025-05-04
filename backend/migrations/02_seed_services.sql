-- Insert services
INSERT INTO services (
    name, 
    short_description, 
    description, 
    price, 
    discounted_price, 
    is_subscription, 
    subscription_interval, 
    image, 
    is_active
) VALUES (
    'Hälsokontroll - Kvinna',
    'Omfattande hälsokontroll anpassad för kvinnor',
    'Vår omfattande hälsokontroll för kvinnor innehåller alla viktiga hälsokontroller som rekommenderas för kvinnor i alla åldrar. Inkluderar blodprover, gynekologisk undersökning, och personlig konsultation.',
    3495.00,
    2995.00,
    FALSE,
    NULL,
    '/assets/images/health-check-women.jpg',
    TRUE
), (
    'Hälsokontroll - Man',
    'Omfattande hälsokontroll anpassad för män',
    'Vår omfattande hälsokontroll för män innehåller alla viktiga hälsokontroller som rekommenderas för män i alla åldrar. Inkluderar blodprover, prostataundersökning, och personlig konsultation.',
    3495.00,
    2995.00,
    FALSE,
    NULL,
    '/assets/images/health-check-men.jpg',
    TRUE
), (
    'Blodprov - Bas',
    'Grundläggande blodprover för allmän hälsokontroll',
    'Vårt baspaket för blodprover ger dig en bra grund för att kontrollera din allmänna hälsostatus. Perfekt för regelbundna kontroller.',
    995.00,
    NULL,
    FALSE,
    NULL,
    '/assets/images/blood-test-basic.jpg',
    TRUE
), (
    'Blodprov - Premium',
    'Omfattande blodanalys med över 20 parametrar',
    'Vårt premiumpaket för blodprover ger en omfattande analys av din hälsa med över 20 olika parametrar. Inkluderar personlig läkarkonsultation för genomgång av resultaten.',
    2495.00,
    1995.00,
    FALSE,
    NULL,
    '/assets/images/blood-test-premium.jpg',
    TRUE
), (
    'Blodprov - Prenumeration',
    'Regelbundna blodprover för kontinuerlig hälsokontroll',
    'Vår prenumerationstjänst för blodprover ger dig möjlighet att regelbundet kontrollera din hälsa. Perfekt för dig som vill följa din hälsoutveckling över tid.',
    695.00,
    NULL,
    TRUE,
    'Var 3:e månad',
    '/assets/images/blood-test-subscription.jpg',
    TRUE
);

-- Insert service features
INSERT INTO service_features (service_id, feature) VALUES 
-- Hälsokontroll - Kvinna (ID: 1)
(1, 'Komplett blodprov (inkl. hormoner)'),
(1, 'Gynekologisk undersökning'),
(1, 'Bröstundersökning'),
(1, 'BMI och kroppssammansättning'),
(1, 'Blodtryck och hjärthälsa'),
(1, 'Läkarkonsultation med genomgång av resultat'),
(1, 'Personlig hälsoplan'),

-- Hälsokontroll - Man (ID: 2)
(2, 'Komplett blodprov (inkl. hormoner)'),
(2, 'Prostataundersökning (PSA)'),
(2, 'BMI och kroppssammansättning'),
(2, 'Blodtryck och hjärthälsa'),
(2, 'Läkarkonsultation med genomgång av resultat'),
(2, 'Personlig hälsoplan'),

-- Blodprov - Bas (ID: 3)
(3, 'Hemoglobin'),
(3, 'Blodsocker'),
(3, 'Kolesterol'),
(3, 'Lever- och njurvärden'),
(3, 'Digitala resultat inom 2 arbetsdagar'),
(3, 'Kort tolkning av resultaten'),

-- Blodprov - Premium (ID: 4)
(4, 'Över 20 olika blodparametrar'),
(4, 'Hormonnivåer'),
(4, 'Vitaminer och mineraler'),
(4, 'Immunfunktion'),
(4, 'Inflammationsmarkörer'),
(4, 'Läkarkonsultation för genomgång av resultat'),
(4, 'Personlig hälsoplan'),

-- Blodprov - Prenumeration (ID: 5)
(5, 'Bas-blodprov var 3:e månad'),
(5, 'Personlig hälsotrend över tid'),
(5, 'Digital resultatgenomgång'),
(5, 'Prioriterad bokning'),
(5, 'Förmånligt prenumerationspris');