CREATE TABLE Promotions (
                            PromotionCode VARCHAR(20) PRIMARY KEY,
                            Description TEXT,
                            StartDate DATE,
                            EndDate DATE,
                            DiscountType ENUM('Percentage', 'FixedValue'),
                            DiscountValue DECIMAL(10, 2),
                            CompanyCode VARCHAR(20)
);


CREATE TABLE PromotionUsage (
                                UsageID INT PRIMARY KEY,
                                PromotionCode VARCHAR(20),
                                CustomerID INT,
                                TransactionID INT,
                                UsageDate TIMESTAMP,
                                FOREIGN KEY (PromotionCode) REFERENCES Promotions(PromotionCode),
);
