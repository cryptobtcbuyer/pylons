// Mocks generated by Mockito 5.3.2 from annotations
// in pylons_wallet/test/waleed_test.dart.
// Do not manually edit this file.

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:ui' as _i4;

import 'package:mockito/mockito.dart' as _i1;
import 'package:pylons_wallet/pages/presenting_onboard_page/viewmodel/accept_policy_viewmodel.dart'
    as _i3;
import 'package:pylons_wallet/services/repository/repository.dart' as _i2;

// ignore_for_file: type=lint
// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types
// ignore_for_file: subtype_of_sealed_class

class _FakeRepository_0 extends _i1.SmartFake implements _i2.Repository {
  _FakeRepository_0(
    Object parent,
    Invocation parentInvocation,
  ) : super(
          parent,
          parentInvocation,
        );
}

/// A class which mocks [AcceptPolicyViewModel].
///
/// See the documentation for Mockito's code generation for more information.
class MockAcceptPolicyViewModel extends _i1.Mock
    implements _i3.AcceptPolicyViewModel {
  @override
  _i2.Repository get repository => (super.noSuchMethod(
        Invocation.getter(#repository),
        returnValue: _FakeRepository_0(
          this,
          Invocation.getter(#repository),
        ),
        returnValueForMissingStub: _FakeRepository_0(
          this,
          Invocation.getter(#repository),
        ),
      ) as _i2.Repository);
  @override
  bool get isCheckTermServices => (super.noSuchMethod(
        Invocation.getter(#isCheckTermServices),
        returnValue: false,
        returnValueForMissingStub: false,
      ) as bool);
  @override
  set isCheckTermServices(bool? _isCheckTermServices) => super.noSuchMethod(
        Invocation.setter(
          #isCheckTermServices,
          _isCheckTermServices,
        ),
        returnValueForMissingStub: null,
      );
  @override
  bool get isCheckPrivacyPolicy => (super.noSuchMethod(
        Invocation.getter(#isCheckPrivacyPolicy),
        returnValue: false,
        returnValueForMissingStub: false,
      ) as bool);
  @override
  set isCheckPrivacyPolicy(bool? _isCheckPrivacyPolicy) => super.noSuchMethod(
        Invocation.setter(
          #isCheckPrivacyPolicy,
          _isCheckPrivacyPolicy,
        ),
        returnValueForMissingStub: null,
      );
  @override
  bool get hasListeners => (super.noSuchMethod(
        Invocation.getter(#hasListeners),
        returnValue: false,
        returnValueForMissingStub: false,
      ) as bool);
  @override
  void toggleCheckTermServices(bool? value) => super.noSuchMethod(
        Invocation.method(
          #toggleCheckTermServices,
          [value],
        ),
        returnValueForMissingStub: null,
      );
  @override
  void toggleCheckPrivacyPolicy(bool? value) => super.noSuchMethod(
        Invocation.method(
          #toggleCheckPrivacyPolicy,
          [value],
        ),
        returnValueForMissingStub: null,
      );
  @override
  void setUserAcceptPolicies() => super.noSuchMethod(
        Invocation.method(
          #setUserAcceptPolicies,
          [],
        ),
        returnValueForMissingStub: null,
      );
  @override
  bool getUserAcceptPolicies() => (super.noSuchMethod(
        Invocation.method(
          #getUserAcceptPolicies,
          [],
        ),
        returnValue: false,
        returnValueForMissingStub: false,
      ) as bool);
  @override
  void addListener(_i4.VoidCallback? listener) => super.noSuchMethod(
        Invocation.method(
          #addListener,
          [listener],
        ),
        returnValueForMissingStub: null,
      );
  @override
  void removeListener(_i4.VoidCallback? listener) => super.noSuchMethod(
        Invocation.method(
          #removeListener,
          [listener],
        ),
        returnValueForMissingStub: null,
      );
  @override
  void dispose() => super.noSuchMethod(
        Invocation.method(
          #dispose,
          [],
        ),
        returnValueForMissingStub: null,
      );
  @override
  void notifyListeners() => super.noSuchMethod(
        Invocation.method(
          #notifyListeners,
          [],
        ),
        returnValueForMissingStub: null,
      );
}
