"""Tests for pipdeps script."""

import mock
import os
import unittest


import pipdeps
import utility


class TestGetParser(unittest.TestCase):

    def test_list_defaults(self):
        parser = pipdeps.get_parser("pipdeps.py")
        args = parser.parse_args(["list"])
        self.assertEqual("list", args.command)
        self.assertEqual("~/cloud-city", args.cloud_city)
        self.assertEqual(False, args.verbose)

    def test_install_verbose(self):
        parser = pipdeps.get_parser("pipdeps.py")
        args = parser.parse_args(["-v", "install"])
        self.assertEqual("install", args.command)
        self.assertEqual(True, args.verbose)

    def test_update_cloud_city(self):
        parser = pipdeps.get_parser("pipdeps.py")
        args = parser.parse_args(["--cloud-city", "/tmp/cc", "update"])
        self.assertEqual("update", args.command)
        self.assertEqual("/tmp/cc", args.cloud_city)

    def test_delete(self):
        parser = pipdeps.get_parser("pipdeps.py")
        args = parser.parse_args(["delete"])
        self.assertEqual("delete", args.command)


class TestS3Connection(unittest.TestCase):

    def test_anon(self):
        sentinel = object()
        with mock.patch("boto.s3.connection.S3Connection", autospec=True,
                        return_value=sentinel) as s3_mock:
            s3 = pipdeps.s3_anon()
            self.assertIs(s3, sentinel)
        s3_mock.assert_called_once_with(anon=True)

    def test_auth_with_rc(self):
        with utility.temp_dir() as fake_city:
            with open(os.path.join(fake_city, "ec2rc"), "wb") as f:
                f.write(
                    "AWS_SECRET_KEY=secret\n"
                    "AWS_ACCESS_KEY=access\n"
                    "export AWS_SECRET_KEY AWS_ACCESS_KEY\n"
                )
            sentinel = object()
            with mock.patch("boto.s3.connection.S3Connection", autospec=True,
                            return_value=sentinel) as s3_mock:
                s3 = pipdeps.s3_auth_with_rc(fake_city)
                self.assertIs(s3, sentinel)
        s3_mock.assert_called_once_with("access", "secret")


class TestRunPipInstall(unittest.TestCase):

    req_path = os.path.join(os.path.dirname(__file__), "requirements.txt")

    def test_added_args(self):
        with mock.patch("subprocess.check_call", autospec=True) as cc_mock:
            pipdeps.run_pip_install(["--user"])
        cc_mock.assert_called_once_with([
            "pip", "-q", "install", "-r", self.req_path, "--user"])

    def test_verbose(self):
        with mock.patch("subprocess.check_call", autospec=True) as cc_mock:
            pipdeps.run_pip_install(["--download", "/tmp/pip"], verbose=True)
        cc_mock.assert_called_once_with([
            "pip", "install", "-r", self.req_path, "--download", "/tmp/pip"])
