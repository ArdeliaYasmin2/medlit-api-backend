## Medlit API Backend

How to use the API:

-   Open your browser and go to endpoint https://medlit-app.et.r.appspot.com/api/medlit/
-   You should see json that returns

    ```
    {
        "error": "false",
    	"message": "Welcome to Medlit API",
    }
    ```

-   You can use the following endpoints to test the API:
-   Endpoint: https://medlit-app.et.r.appspot.com/api/medlit/

-   GET Method

    -   `/medicine/get/all`: returns all medicines.

        ```
        {
            "error": "false",
            "medicineList": [
                {
                    "id": 1,
                    "generic_name": "Amlodipine Besilate",
                    "photo_url": "https://storage.googleapis.commedlit-bucket/amlodipine.jpg",
                    "purpose": "Treat high blood pressure, certain types of angina and coronary artery disease",
                    "side_effects": "Swelling of the hands, feet, ankles, or lower legs, headache, upset stomach, nausea, stomach pain, dizziness or lightheadedness, drowsiness, etc",
                    "contraindication": "Hypersensitivity",
                    "dosage": "5 to 10 mg once a day for Usual Adult Dose, 2.5 to 5 mg once a day for Usual Pediatic Dose",
                    "ingredients": "Amlodipine 5 mg",
                    "created_at": "2022-06-12T10:38:52.693161Z",
                    "updated_at": "2022-06-12T10:38:52.693161Z"
                },
                {
                    "id": 2,
                    "generic_name": "Antasida Doen",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/antasida-doen.jpg",
                    "purpose": "Relieve heartburn, sour stomach, or acid indigestion",
                    "side_effects": "Diarrhea, constipation, fatigue without apparent cause, loss ofappetite, dehydration, etc",
                    "contraindication": "Severe Liver disfunction",
                    "dosage": "1 to 2 tablets 3 to 4 times a day for Usual Adult Dose, 0.5 to 1 tablet 3 to 4 times a day for Usual Pediatric Dose",
                    "ingredients": "Alumunium Hydroxide 200 mg, Magnesium Hydroxide 200 mg",
                    "created_at": "2022-06-12T10:39:17.345561Z",
                    "updated_at": "2022-06-12T10:39:17.345561Z"
                },
                {
                    "id": 3,
                    "generic_name": "Incidal",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/incidal.jpg",
                    "purpose": "Relieve the symptoms of allergic inflammation of the nasal airways due to allergens",
                    "side_effects": "Headache, dizziness, sleepiness, tiredness",
                    "contraindication": "Hypersensitivity,Severe Kidney Disorders",
                    "dosage": "5 to 10 mg once a day for Usual Adult Dose, 2.5 to 5 mg once a day for Usual Pediatric Dose",
                    "ingredients": "Cetirizine 10 mg",
                    "created_at": "2022-06-12T10:39:39.060857Z",
                    "updated_at": "2022-06-12T10:39:39.060857Z"
                },
                {
                    "id": 4,
                    "generic_name": "Mefenamic Acid",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/mefenamid-acid.jpg",
                    "purpose": "Relieve mild to moderate pain, including menstrual pain (pain that happens before or during a menstrual period)",
                    "side_effects": "Diarrhea, constipation, gas or bloating, headache, dizziness, nervousness, ringing in the ears, etc",
                    "contraindication": "Hypersensitivity reaction to mefenamic acid. Patients taking aspirin experience bronchospasm, allergic rhinitis and urticaria. Patients with gastric ulcers or inflammation of the digestive tract. Patients with kidney disorders",
                    "dosage": "500 mg once, then 250 mg every 6 hours as needed for Usual Adult Dose",
                    "ingredients": "Mefenamic Acid 250 mg",
                    "created_at": "2022-06-12T10:39:58.083735Z",
                    "updated_at": "2022-06-12T10:39:58.083735Z"
                },
                {
                    "id": 5,
                    "generic_name": "Paracetamol",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/paracetamol.jpg",
                    "purpose": "Pain reliever and a fever reducer that can be used to treat many conditions such as headache, muscle aches, arthritis, backache, toothaches, colds, and fevers",
                    "side_effects": "Low fever with nausea, stomach pain, loss of appetite, dark urine, clay-colored stools, jaundice (yellowing of the skin or eyes), etc",
                    "contraindication": "Hypersensitivity to acetaminophen, severe hepatic impairment, or severe active hepatic disease",
                    "dosage": "500 to 1000 mg every 4 to 6 hours for Usual Adult Dose, 325 to 650 mg every 4 to 6 hours for Usual Pediatric Dose (>= 12 years)",
                    "ingredients": "Paracetamol 500 mg",
                    "created_at": "2022-06-12T10:40:20.545594Z",
                    "updated_at": "2022-06-12T10:40:20.545594Z"
                },
                {
                    "id": 6,
                    "generic_name": "Simvastatin",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/simvastatin.jpg",
                    "purpose": "Treat high cholesterol and to reduce risk of heart disease",
                    "side_effects": "Unexplained muscle weakness or tenderness, persistent muscle pain, abdominal pain, fever, dark-colored urine, etc",
                    "contraindication": "Liver Function Failure, Hypersensitivity, pregnant & lactating women",
                    "dosage": "10 to 20 mg once a day, then 5 to 40 mg once a day for Maintenance Dose (Only Usual Adult Dose)",
                    "ingredients": "Simvastatin 10 mg",
                    "created_at": "2022-06-12T10:40:41.521459Z",
                    "updated_at": "2022-06-12T10:40:41.521459Z"
                }
            ],
            "message": "Medicine list"

        }
        ```

    -   `/medicine/get/{id}`: returns a medicine with the given id.

        ```
        {
            "error": "false",
            "medicine": {
                "id": 1,
                "generic_name": "Amlodipine Besilate",
                "photo_url": "https://storage.googleapis.com/medlit-bucket/amlodipine.jpg",
                "purpose": "Treat high blood pressure, certain types of angina and coronary artery disease",
                "side_effects": "Swelling of the hands, feet, ankles, or lower legs, headache, upset stomach, nausea, stomach pain, dizziness or lightheadedness, drowsiness, etc",
                "contraindication": "Hypersensitivity",
                "dosage": "5 to 10 mg once a day for Usual Adult Dose, 2.5 to 5 mg once a day for Usual Pediatic Dose",
                "ingredients": "Amlodipine 5 mg",
                "created_at": "2022-06-12T10:38:52.693161Z",
                "updated_at": "2022-06-12T10:38:52.693161Z"
            },
            "message": "Medicine found"
        }
        ```

    -   `/medicine/search?generic_name={name}`: returns all medicines with the given name.

        ```
        {
            "error": "false",
            "medicineList": [
                {
                    "id": 6,
                    "generic_name": "Simvastatin",
                    "photo_url": "https://storage.googleapis.com/medlit-bucket/simvastatin.jpg",
                    "purpose": "Treat high cholesterol and to reduce risk of heart disease",
                    "side_effects": "Unexplained muscle weakness or tenderness, persistent muscle pain, abdominal pain, fever, dark-colored urine, etc",
                    "contraindication": "Liver Function Failure, Hypersensitivity, pregnant & lactating women",
                    "dosage": "10 to 20 mg once a day, then 5 to 40 mg once a day for Maintenance Dose (Only Usual Adult Dose)",
                    "ingredients": "Simvastatin 10 mg",
                    "created_at": "2022-06-12T10:40:41.521459Z",
                    "updated_at": "2022-06-12T10:40:41.521459Z"
                }
            ],
            "message": "Medicine found"
        }
        ```

-   POST Method

    -   `/medicine/add`: adds a new medicine.

        ```
        {
            "error": "false",
            "message": "Medicine created"
        }
        ```

    -   `/register`: registers a new user.

        ```
        {
            "error": "false",
            "message": "Registration Successful"
        }
        ```

    -   `/login`: logs in a user.

        ```
        {
            "error": "false",
            "message": "Login Successful"
        }
        ```

<br>

Thanks for using Medlit!

Made with ❤ by MEDLIT Team

Copyright © 2022 MEDLIT. All rights reserved.
