CREATE TABLE Promotions (
                            PromotionCode VARCHAR(20) PRIMARY KEY,
                            Description TEXT,
                            StartDate DATE,
                            EndDate DATE,
                            DiscountType ENUM('Percentage', 'FixedValue'),
                            DiscountValue DECIMAL(10, 2),
                            CompanyCode VARCHAR(20),
                            Status VARCHAR(12) CHECK
                                (CASE WHEN EXTRACT(EPOCH FROM StartDate) > EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) THEN 'active'  END IN ('active', 'inactive'))

);


CREATE TABLE PromotionUsage (
                                UsageID INT PRIMARY KEY,
                                PromotionCode VARCHAR(20),
                                CustomerID INT,
                                TransactionID INT,
                                UsageDate TIMESTAMP,
                                FOREIGN KEY (PromotionCode) REFERENCES Promotions(PromotionCode),
);


start_time TIMESTAMP,
                            status VARCHAR(12) CHECK
                                (CASE WHEN EXTRACT(EPOCH FROM StartDate) < EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) THEN 'inactive' ELSE 'active' END IN ('active', 'inactive'))