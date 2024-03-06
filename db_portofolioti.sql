-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 06, 2024 at 06:58 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_portofolioti`
--

-- --------------------------------------------------------

--
-- Table structure for table `aplikasi`
--

CREATE TABLE `aplikasi` (
  `id_aplikasi` int(50) NOT NULL,
  `nama_aplikasi` varchar(200) NOT NULL,
  `tahun` year(4) NOT NULL,
  `alias` varchar(200) NOT NULL,
  `url` varchar(200) NOT NULL,
  `bussinesowner` varchar(200) NOT NULL,
  `Acquisition` varchar(200) NOT NULL,
  `fe_language` varchar(200) NOT NULL,
  `fe_framework` varchar(200) NOT NULL,
  `be_language` varchar(200) NOT NULL,
  `be_framework` varchar(200) NOT NULL,
  `id_perusahaan` int(50) NOT NULL,
  `id_category` int(50) NOT NULL,
  `id_status` int(50) NOT NULL,
  `id_valuechain` int(50) NOT NULL,
  `id_provider` int(50) NOT NULL,
  `id_asset` int(50) NOT NULL,
  `id_contract` int(50) NOT NULL,
  `id_vendor` int(50) NOT NULL,
  `id_method` int(50) NOT NULL,
  `id_type` int(50) NOT NULL,
  `id_architecture` int(50) NOT NULL,
  `id_infrastructure` int(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `architectures`
--

CREATE TABLE `architectures` (
  `id_architecture` int(11) NOT NULL,
  `nama_architecture` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `assets`
--

CREATE TABLE `assets` (
  `id_asset` int(11) NOT NULL,
  `nama_asset` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id_category` int(50) NOT NULL,
  `nama_category` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `contracts`
--

CREATE TABLE `contracts` (
  `id_contract` int(50) NOT NULL,
  `nama_contract` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `infrastructures`
--

CREATE TABLE `infrastructures` (
  `id_infrastructure` int(11) NOT NULL,
  `nama_infrastucture` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `methods`
--

CREATE TABLE `methods` (
  `id_method` int(50) NOT NULL,
  `nama_method` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `moduls`
--

CREATE TABLE `moduls` (
  `id_modul` int(50) NOT NULL,
  `nama_modul` varchar(200) NOT NULL,
  `id_aplikasi` int(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `perusahaan`
--

CREATE TABLE `perusahaan` (
  `id_perusahaan` int(50) NOT NULL,
  `nama_perusahaan` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `providers`
--

CREATE TABLE `providers` (
  `id_provider` int(50) NOT NULL,
  `nama_provider` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `status`
--

CREATE TABLE `status` (
  `id_status` int(50) NOT NULL,
  `nama_status` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `types`
--

CREATE TABLE `types` (
  `id_type` int(50) NOT NULL,
  `nama_type` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(200) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` text NOT NULL,
  `nama` varchar(50) NOT NULL,
  `role` int(11) NOT NULL,
  `created_at` decimal(65,0) NOT NULL,
  `updated_at` decimal(65,0) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `nama`, `role`, `created_at`, `updated_at`) VALUES
('88f5596f-2128-4df3-bd7c-e9546db5b10c', 'bintangf', 'bintangfajar@gmail.com', '$2a$10$A0gBXeMg6IOqiCyPAafPiOxtWpSt.AVLxgHCPVu6GEMvIIpSeKOUi', 'Bintang Fajar', 1, '1709659342', '1709659342'),
('d994b321-c72f-4f10-8e0d-69547f84b9c2', 'bintang', 'bintangfazr@gmail.com', '$2a$10$olafcNWsxD7pW4hi4saLmuiREpzOwR4fgnBcXCQNpCRmzM3cfwmjK', 'Bintang Fajar', 1, '1709657448', '1709657448');

-- --------------------------------------------------------

--
-- Table structure for table `valuechains`
--

CREATE TABLE `valuechains` (
  `id_valuechain` int(50) NOT NULL,
  `nama_valuechain` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `vendors`
--

CREATE TABLE `vendors` (
  `id_vendor` int(50) NOT NULL,
  `nama_vendor` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `aplikasi`
--
ALTER TABLE `aplikasi`
  ADD PRIMARY KEY (`id_aplikasi`),
  ADD KEY `id_architecture` (`id_architecture`),
  ADD KEY `id_asset` (`id_asset`),
  ADD KEY `id_category` (`id_category`),
  ADD KEY `id_contract` (`id_contract`),
  ADD KEY `id_infrastructure` (`id_infrastructure`),
  ADD KEY `id_method` (`id_method`),
  ADD KEY `id_perusahaan` (`id_perusahaan`),
  ADD KEY `id_provider` (`id_provider`),
  ADD KEY `id_status` (`id_status`),
  ADD KEY `id_type` (`id_type`),
  ADD KEY `id_valuechain` (`id_valuechain`),
  ADD KEY `id_vendor` (`id_vendor`);

--
-- Indexes for table `architectures`
--
ALTER TABLE `architectures`
  ADD PRIMARY KEY (`id_architecture`);

--
-- Indexes for table `assets`
--
ALTER TABLE `assets`
  ADD PRIMARY KEY (`id_asset`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id_category`);

--
-- Indexes for table `contracts`
--
ALTER TABLE `contracts`
  ADD PRIMARY KEY (`id_contract`);

--
-- Indexes for table `infrastructures`
--
ALTER TABLE `infrastructures`
  ADD PRIMARY KEY (`id_infrastructure`);

--
-- Indexes for table `methods`
--
ALTER TABLE `methods`
  ADD PRIMARY KEY (`id_method`);

--
-- Indexes for table `moduls`
--
ALTER TABLE `moduls`
  ADD PRIMARY KEY (`id_modul`),
  ADD KEY `id_aplikasi` (`id_aplikasi`);

--
-- Indexes for table `perusahaan`
--
ALTER TABLE `perusahaan`
  ADD PRIMARY KEY (`id_perusahaan`);

--
-- Indexes for table `providers`
--
ALTER TABLE `providers`
  ADD PRIMARY KEY (`id_provider`);

--
-- Indexes for table `status`
--
ALTER TABLE `status`
  ADD PRIMARY KEY (`id_status`);

--
-- Indexes for table `types`
--
ALTER TABLE `types`
  ADD PRIMARY KEY (`id_type`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `valuechains`
--
ALTER TABLE `valuechains`
  ADD PRIMARY KEY (`id_valuechain`);

--
-- Indexes for table `vendors`
--
ALTER TABLE `vendors`
  ADD PRIMARY KEY (`id_vendor`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `aplikasi`
--
ALTER TABLE `aplikasi`
  MODIFY `id_aplikasi` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `architectures`
--
ALTER TABLE `architectures`
  MODIFY `id_architecture` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `assets`
--
ALTER TABLE `assets`
  MODIFY `id_asset` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id_category` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `contracts`
--
ALTER TABLE `contracts`
  MODIFY `id_contract` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `infrastructures`
--
ALTER TABLE `infrastructures`
  MODIFY `id_infrastructure` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `methods`
--
ALTER TABLE `methods`
  MODIFY `id_method` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `moduls`
--
ALTER TABLE `moduls`
  MODIFY `id_modul` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `perusahaan`
--
ALTER TABLE `perusahaan`
  MODIFY `id_perusahaan` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `providers`
--
ALTER TABLE `providers`
  MODIFY `id_provider` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `status`
--
ALTER TABLE `status`
  MODIFY `id_status` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `types`
--
ALTER TABLE `types`
  MODIFY `id_type` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `valuechains`
--
ALTER TABLE `valuechains`
  MODIFY `id_valuechain` int(50) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `vendors`
--
ALTER TABLE `vendors`
  MODIFY `id_vendor` int(50) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `aplikasi`
--
ALTER TABLE `aplikasi`
  ADD CONSTRAINT `aplikasi_ibfk_1` FOREIGN KEY (`id_architecture`) REFERENCES `architectures` (`id_architecture`),
  ADD CONSTRAINT `aplikasi_ibfk_10` FOREIGN KEY (`id_type`) REFERENCES `types` (`id_type`),
  ADD CONSTRAINT `aplikasi_ibfk_11` FOREIGN KEY (`id_valuechain`) REFERENCES `valuechains` (`id_valuechain`),
  ADD CONSTRAINT `aplikasi_ibfk_12` FOREIGN KEY (`id_vendor`) REFERENCES `vendors` (`id_vendor`),
  ADD CONSTRAINT `aplikasi_ibfk_2` FOREIGN KEY (`id_asset`) REFERENCES `assets` (`id_asset`),
  ADD CONSTRAINT `aplikasi_ibfk_3` FOREIGN KEY (`id_category`) REFERENCES `categories` (`id_category`),
  ADD CONSTRAINT `aplikasi_ibfk_4` FOREIGN KEY (`id_contract`) REFERENCES `contracts` (`id_contract`),
  ADD CONSTRAINT `aplikasi_ibfk_5` FOREIGN KEY (`id_infrastructure`) REFERENCES `infrastructures` (`id_infrastructure`),
  ADD CONSTRAINT `aplikasi_ibfk_6` FOREIGN KEY (`id_method`) REFERENCES `methods` (`id_method`),
  ADD CONSTRAINT `aplikasi_ibfk_7` FOREIGN KEY (`id_perusahaan`) REFERENCES `perusahaan` (`id_perusahaan`),
  ADD CONSTRAINT `aplikasi_ibfk_8` FOREIGN KEY (`id_provider`) REFERENCES `providers` (`id_provider`),
  ADD CONSTRAINT `aplikasi_ibfk_9` FOREIGN KEY (`id_status`) REFERENCES `status` (`id_status`);

--
-- Constraints for table `moduls`
--
ALTER TABLE `moduls`
  ADD CONSTRAINT `moduls_ibfk_1` FOREIGN KEY (`id_aplikasi`) REFERENCES `aplikasi` (`id_aplikasi`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
